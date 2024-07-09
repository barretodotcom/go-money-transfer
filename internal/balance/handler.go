package balance

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func HandleTransfer(ctx *fiber.Ctx) error {
	var transferRequest TransferRequest
	err := ctx.BodyParser(&transferRequest)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	debtorId := ctx.Locals("balanceId")
	transferRequest.DebtorID = debtorId.(string)
	balanceService := BuildBalanceService()

	err = balanceService.Transfer(transferRequest)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(http.StatusOK).JSON(nil)
}
