package user

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HandleCreateUser(c *fiber.Ctx) error {
	var createUserDto CreateUser
	err := c.BodyParser(&createUserDto)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	service := BuildUserService()
	err = service.CreateUser(createUserDto)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(http.StatusCreated)
}
