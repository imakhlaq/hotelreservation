package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/imakhlaq/hotelreservation/db"
	"github.com/imakhlaq/hotelreservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserHandler struct {
	userStore db.UserStore //userStore is interface so u can embed any struct that satisfy the interface
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}
func (u UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	//before deleting check the if its exits
	_, err := u.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]any{"massage": "Record Not found"})
		}
	}

	if err := u.userStore.DeleteUser(c.Context(), id); err != nil {
		return err
	}
	//Sprintf is for string format Printf is for writing to the terminal.
	res := fmt.Sprintf("User with id %s is deleted", id)
	return c.JSON(map[string]any{"message": res})
}

func (u UserHandler) HandlePostUser(c *fiber.Ctx) error {
	//parsing the body
	var params types.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	//validating params
	if err := params.Validate(); len(err) != 0 {
		return c.Status(400).JSON(err)
	}
	//creaing user type from parms
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	//insert in the db
	insertedUser, err := u.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}

func (u UserHandler) HandleGetUser(c *fiber.Ctx) error {
	id := c.Params("id") // id from req param

	user, err := u.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"message": "User does't exit."})
		}
		return err
	}
	return c.JSON(user)
}
func (u UserHandler) HandleUsers(c *fiber.Ctx) error {
	users, err := u.userStore.GetAllUsers(c.Context())
	if err != nil {
		return fmt.Errorf("NO USERS IN DB")
	}
	return c.JSON(users)
}
func (u UserHandler) HandleUpdate(c *fiber.Ctx) error {
	var (
		id    = c.Params("id")
		param map[string]any
	)
	if err := c.BodyParser(&param); err != nil {
		return err
	}
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": oid}

	if err = u.userStore.UpdateUser(c.Context(), filter, param); err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": id})
}
