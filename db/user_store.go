package db

import "github.com/imakhlaq/hotelreservation/types"

//whatever you are using need to provide this interface implementation
type UserStore interface {
	GetUserByID(string) (*types.User, error)
}

//example
type MongoUserStore struct {
}

// if u want to use postgres
// type PostgresUserStore struct {
// }
