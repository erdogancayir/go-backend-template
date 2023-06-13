package controller

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/badoux/checkmail"
	"github.com/erdogancayir/nargileapi/bootstrap"
	"github.com/erdogancayir/nargileapi/domain"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
	Db            *mongo.Database
}

//Hnadler func has a cho.Context param. This context obj contains from req and res info's
func (sc *SignupController) Signup(c echo.Context) error { // <--
	var request domain.SignupRequest
	//Bind func tries to the incoming HTTP request to the request variable.

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()}) // <--
	}

	log.Println(request.Email)
	_err := checkmail.ValidateFormat(request.Email)
	if _err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Geçerli bir e-posta adresi girin"})
	}

	mailsCollection := sc.Db.Collection("mails")
	userCode := request.MailCode

	stdContext := c.Request().Context()

	// Check if the code from the request matches the one in the mails collection
	var mail domain.MailSave
	mailFilter := bson.D{{Key: "email", Value: request.Email}, {Key: "code", Value: userCode}}
	mailErr := mailsCollection.FindOne(stdContext, mailFilter).Decode(&mail)
	if mailErr != nil {
		if errors.Is(mailErr, mongo.ErrNoDocuments) {
			return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Invalid email or code"})
		}
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: mailErr.Error()})
	}

	_, __err := sc.SignupUsecase.GetUserByEmail(stdContext, request.Email)
	if __err == nil {
		return c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists with the given email"})
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()}) // <--
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:             primitive.NewObjectID(),
		Name:           request.Name,
		SurName:        request.SurName,
		Email:          request.Email,
		Password:       request.Password,
		PhoneNumber:    request.PhoneNumber,
		NickName:       request.NickName,
		Addresses:      request.Addresses,
		Gender:         request.Gender,
		DateOfBirth:    request.DateOfBirth,
		ProfilePicture: request.ProfilePicture,
		TwoFactorAuth:  request.TwoFactorAuth,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	/* log.Fatal(user) */

	err = sc.SignupUsecase.Create(stdContext, &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()}) // <--
	}

	accessToken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()}) // <--
	}

	refreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()}) // <--
	}

	signupResponse := domain.SignupResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return c.JSON(http.StatusOK, signupResponse) // <--
}
