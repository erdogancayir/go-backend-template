package controller

import (
	"net/http"

	"github.com/erdogancayir/nargileapi/bootstrap"
	"github.com/erdogancayir/nargileapi/domain"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MailController struct {
	MailUsecase domain.MailRepository
	Env         *bootstrap.Env
	Db          *mongo.Database
}

func (sc *MailController) SendMail(c echo.Context) error {
	var request domain.MailRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
	}

	verificationCode, err := sc.MailUsecase.SendMail(request.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	collection := sc.Db.Collection("mails")
	filter := bson.M{"email": request.Email}
	update := bson.M{
		"$set": bson.M{
			"email": request.Email,
			"code":  verificationCode,
		},
	}

	updateOpts := options.Update().SetUpsert(true)

	_, err = collection.UpdateOne(c.Request().Context(), filter, update, updateOpts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: "Failed to upsert mail: " + err.Error()})
	}
	return c.JSON(http.StatusOK, domain.MailResponse{Message: "SuccessMail"})
}
