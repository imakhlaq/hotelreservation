package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	bcryptCost = 12
)

//bson:"_id" => what should be db colum name
//json:"id" => for json conversion
//json:"id,omitempty" => ignore if field is empty
//json:"-" => ignore field while JSON marshaling
//bson:"_id,omitempty =>Because we want mongodb to create the id for us

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Username          string             `bson:"user_name" json:"username"`
	EncriptedPassword string             `bson:"encripted_password" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {

	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Username:          params.Username,
		EncriptedPassword: string(encpw),
	}, nil
}
