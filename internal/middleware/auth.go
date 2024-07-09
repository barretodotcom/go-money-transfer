package middleware

import (
	"net/http"
	"strings"

	"github.com/go-money-transfer/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, ok := c.GetReqHeaders()["Authorization"]
		if !ok {
			return c.SendStatus(http.StatusUnauthorized)
		}
		token = strings.Split(token[0], "Bearer ")
		tokenData, err := jwt.ParseToken(token[1])
		if err != nil {
			return c.SendStatus(http.StatusUnauthorized)
		}
		c.Locals("userId", tokenData.UserId)
		c.Locals("balanceId", tokenData.BalanceId)
		return c.Next()
	}
}
