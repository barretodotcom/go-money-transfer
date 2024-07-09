package auth

import (
	"net/http"

	"github.com/go-money-transfer/internal/balance"
	"github.com/go-money-transfer/internal/database"
	"github.com/go-money-transfer/internal/user"
	"github.com/gofiber/fiber/v2"
)

func HandleLogin(ctx *fiber.Ctx) error {
	var login Login
	err := ctx.BodyParser(&login)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	userRepository := user.UserRepository{DB: database.GetConnection()}
	balanceRepository := balance.BalanceRepository{DB: database.GetConnection()}
	service := AuthService{UsersRepository: &userRepository, BalanceRepository: &balanceRepository}

	token, err := service.AuthUser(login)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{"token": token})
}
