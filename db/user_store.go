package db

import (
	"context"

	"github.com/imakhlaq/hotelreservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userColl = "users"
)

// UserStore needed to be implemented by any DB u  want to use
type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetAllUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
}

// MongoUserStore it uses mongoose
type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: client,
		coll:   client.Database(DBNAME).Collection(userColl),
	}
}

func (m MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {

	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var user types.User
	if err = m.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (m MongoUserStore) GetAllUsers(ctx context.Context) ([]*types.User, error) {

	curr, err := m.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var users []*types.User
	//-------curr.Decode(ctx, &users); for one
	if err = curr.All(ctx, &users); err != nil { //for all
		return nil, err
	}
	return users, nil
}

func (m MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := m.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

// example
// if u want to use postgres
// type PostgresUserStore struct {
// }
