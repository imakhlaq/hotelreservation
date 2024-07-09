package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/imakhlaq/hotelreservation/types"
)

func HandleUser(c *fiber.Ctx) error {

	id := c.Params("id") // id from req param

	return c.JSON(map[string]string{"name": "Akhlaq Ahmad"})
}
func HandleUsers(c *fiber.Ctx) error {
	u := &types.User{
		ID:        "1021039019318",
		FirstName: "Akhlaq",
		LastName:  "Ahmad",
	}
	return c.JSON(u)
}
