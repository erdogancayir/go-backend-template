package domain

import (
	"context"
)

type SignupRequest struct {
	Name           string    `form:"name" binding:"required"`
	SurName        string    `form:"surname" binding:"required"`
	Email          string    `form:"email" binding:"required,email"`
	Password       string    `form:"password" binding:"required"`
	PhoneNumber    string    `form:"phonenumber" binding:"required"`
	NickName       string    `form:"nickname"`
	Addresses      []Address `form:"addresses" binding:"required,dive"` // Consider using dive for nested validation
	Gender         string    `form:"gender"`                            // Assuming Gender can only be Male or Female
	DateOfBirth    string    `form:"dateofbirth"`
	ProfilePicture string    `form:"profilepicture"`
	TwoFactorAuth  bool      `form:"twofactorauth"`
	SentMail       bool      `form:"sentmail" default:"false"`
	MailCode       string    `form:"mailcode" binding:"required"`
}

type SignupResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type SuccessResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type SignupUsecase interface {
	Create(c context.Context, user *User) error
	GetUserByEmail(c context.Context, email string) (User, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
}
