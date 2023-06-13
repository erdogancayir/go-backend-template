package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type MailRequest struct {
	Email    string `form:"email" binding:"required,email"`
	SentMail bool   `form:"sentmail" default:"false"`
}

type MailSave struct {
	ID    primitive.ObjectID `bson:"_id"`
	Email string             `bson:"email"`
	Code  string             `bson:"code"`
}

type MailResponse struct {
	Message string `json:"message"`
}

type MailRepository interface {
	SendMail(to string) (string, error)
}
