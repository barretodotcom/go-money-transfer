package router

import (
	"github.com/go-money-transfer/internal/auth"
	"github.com/go-money-transfer/internal/balance"
	"github.com/go-money-transfer/internal/middleware"
	"github.com/go-money-transfer/internal/user"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(r *fiber.App) {
	usersGroup := r.Group("users")
	usersGroup.Post("", user.HandleCreateUser)

	authGroup := r.Group("auth")
	authGroup.Post("", auth.HandleLogin)

	balanceGroup := r.Group("balance")
	balanceGroup.Post("/transfer", middleware.JWTProtected(), balance.HandleTransfer)
}
