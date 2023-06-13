package controller

import (
	"net/http"

	"github.com/erdogancayir/nargileapi/bootstrap"
	"github.com/erdogancayir/nargileapi/domain"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c echo.Context) error {
	var request domain.LoginRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()}) // <--
	}

	stdContext := c.Request().Context()
	user, err := lc.LoginUsecase.GetUserByEmail(stdContext, request.Email)
	if err != nil {
		return c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: "User not found with the given email"})
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Invalid credentials"})
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
	}

	/* c.SetCookie(&http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	})
	bunlar http ile yapılan cookie get isteklerini engeller, httponly ile
	c.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    refreshToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}) */

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return c.JSON(http.StatusOK, loginResponse)
}
