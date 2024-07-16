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

const (
	minUsername = 4
	minPassword = 7
)

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (p CreateUserParams) Validate() map[string]any {
	err := map[string]any{} //creating and inisilizing empty map

	if len(p.Username) < minUsername {
		err["username"] = "Username is too short"
		//return errors.New("User name is too short")
	}
	if len(p.Password) < minPassword {
		err["password"] = "Password is too short"
		//return fmt.Errorf("Passowrd of %d is too short", len(p.Password))
	}
	return err
}

type User struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username          string             `bson:"user_name" json:"username"`
	EncryptedPassword string             `bson:"encrypted_password" json:"-"`
}

func NewUserFromParams(params CreateUserParams) (*User, error) {

	encpw, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}

	return &User{
		Username:          params.Username,
		EncryptedPassword: string(encpw),
	}, nil
}
