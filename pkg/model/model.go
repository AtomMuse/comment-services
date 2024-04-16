package model

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JwtCustomClaims struct {
	ID           primitive.ObjectID `json:"userID" bson:"userID"`
	Role         string             `json:"role"`
	UserName     string             `json:"username" bson:"username"`
	FirstName    string             `json:"firstname" bson:"firstname"`
	LastName     string             `json:"lastname" bson:"lastname"`
	ProfileImage string             `json:"profile,omitempty" bson:"profile,omitempty"`
	Email        string             `json:"email,omitempty" bson:"email,omitempty"`
	jwt.StandardClaims
}

type User struct {
	ID           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	FirstName    string             `json:"firstname" bson:"firstname"`
	LastName     string             `json:"lastname" bson:"lastname"`
	Username     string             `bson:"username,omitempty" json:"username,omitempty"`
	ProfileImage string             `json:"profile,omitempty" bson:"profile,omitempty"`
}

type UserNoID struct {
	FirstName    string `json:"firstname" bson:"firstname"`
	LastName     string `json:"lastname" bson:"lastname"`
	Username     string `bson:"username,omitempty" json:"username,omitempty"`
	ProfileImage string `json:"profile,omitempty" bson:"profile,omitempty"`
}

type Comment struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" validate:"required"`
	ExhibitionID   primitive.ObjectID `bson:"exhibitionID" json:"exhibitionId" validate:"required"`
	User           User               `bson:"user" json:"user" validate:"required"`
	CommentMessage string             `bson:"commentMessage" json:"commentMessage" validate:"required"`
	CreateDateAt   primitive.DateTime `bson:"createDateAt" json:"createDateAt" validate:"required"`
	UpdateDateAt   primitive.DateTime `bson:"updateDateAt" json:"updateDateAt" validate:"required"`
}

type ResponseComment struct {
	ID             primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty" validate:"required"`
	ExhibitionID   primitive.ObjectID `bson:"exhibitionID" json:"exhibitionId" validate:"required"`
	User           UserNoID           `bson:"user" json:"user" validate:"required"`
	CommentMessage string             `bson:"commentMessage" json:"commentMessage" validate:"required"`
	CreateDateAt   primitive.DateTime `bson:"createDateAt" json:"createDateAt" validate:"required"`
	UpdateDateAt   primitive.DateTime `bson:"updateDateAt,omitempty" json:"updateDateAt,omitempty" `
}

type RequestCreateComment struct {
	ExhibitionID   primitive.ObjectID `bson:"exhibitionID" json:"exhibitionId" validate:"required"`
	User           User               `bson:"user,omitempty" json:"user,omitempty"`
	CommentMessage string             `bson:"commentMessage" json:"commentMessage" validate:"required"`
	UpdateDateAt   primitive.DateTime `bson:"updateDateAt,omitempty" json:"updateDateAt,omitempty" `
	CreateDateAt   primitive.DateTime `bson:"createDateAt,omitempty" json:"createDateAt,omitempty"`
}

type RequestUpdateComment struct {
	CommentMessage string             `bson:"commentMessage" json:"commentMessage" validate:"required"`
	UpdateDateAt   primitive.DateTime `bson:"updateDateAt,omitempty" json:"updateDateAt,omitempty"`
}
