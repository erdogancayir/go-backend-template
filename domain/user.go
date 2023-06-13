package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionUser = "users"
)

type Address struct {
	Street     string `bson:"street"`
	City       string `bson:"city"`
	State      string `bson:"state"`
	Country    string `bson:"country"`
	PostalCode string `bson:"postalcode"`
}

type User struct {
	ID             primitive.ObjectID `bson:"_id"`
	Name           string             `bson:"name"`
	SurName        string             `bson:"surname"`
	Email          string             `bson:"email" json:"email,omitempty"`
	Password       string             `bson:"password"`
	PhoneNumber    string             `bson:"phonenumber"`
	NickName       string             `bson:"nickname"`
	Addresses      []Address          `bson:"addresses"`
	Gender         string             `form:"gender"`
	DateOfBirth    string             `bson:"dateofbirth"`
	CreatedAt      time.Time          `bson:"created_at"`
	UpdatedAt      time.Time          `bson:"updated_at"`
	ProfilePicture string             `bson:"profile_picture"`
	TwoFactorAuth  bool               `bson:"two_factor_auth"`
}

type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context) ([]User, error)
	GetByEmail(c context.Context, email string) (User, error)
	GetByID(c context.Context, id string) (User, error)
}
