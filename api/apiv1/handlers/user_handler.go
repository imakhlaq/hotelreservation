package handlers

import (
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

func (u UserHandler) HandleUser(c *fiber.Ctx) error {

	id := c.Params("id") // id from req param

	user, err := u.userStore.GetUserByID(c.Context(), id)

	if err != nil {
		return c.JSON(map[string]string{"message": "User does't exit."})
	}
	return c.JSON(user)
}
func (u UserHandler) HandleUsers(c *fiber.Ctx) error {
	user := &types.User{
		ID:        "1021039019318",
		FirstName: "Akhlaq",
		LastName:  "Ahmad",
	}
	return c.JSON(user)
}
