package db

import "github.com/imakhlaq/hotelreservation/types"

//whatever you are using need to provide this interface implementation
//u can have multiple databases in production u just need to implement this methods
type UserStore interface {
	GetUserByID(string) (*types.User, error)
}

//example
type MongoUserStore struct {
}

// if u want to use postgres
// type PostgresUserStore struct {
// }
