package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/imakhlaq/hotelreservation/db"
	"github.com/imakhlaq/hotelreservation/types"
)

type UserHandler struct {
	userStore db.UserStore //userStore is interface so u can embed any struct that satisfy the interface
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{userStore: userStore}
}

func (u UserHandler) HandlePostUser(c *fiber.Ctx) error {
	var params types.CreateUserParams
	if err := c.BodyParser(params); err != nil {
		return err
	}
	user, err := types.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := u.userStore.InsertUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}

func (u UserHandler) HandleUser(c *fiber.Ctx) error {
	id := c.Params("id") // id from req param

	user, err := u.userStore.GetUserByID(c.Context(), id)
	if err != nil {
		return c.JSON(map[string]string{"message": "User does't exit."})
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
