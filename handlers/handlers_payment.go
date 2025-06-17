package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"app/models"
	"app/services"
	"gorm.io/gorm"
)

func CreatePayment(db *gorm.DB)fiber.Handler{
	return func(c *fiber.Ctx) error{
		var req models.PaymentRequest
		if err := c.BodyParser(&req); err != nil{
			fmt.Println("BodyParser error:", err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error" : "invalid reuqest body",
			})
		}

		payment, err := services.CreatePaymentService(db, req)
		if err != nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error" : err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"message" : "payment success",
			"payment" : payment,
			"snap_url" : payment.SnapURL,
		})
	}
}