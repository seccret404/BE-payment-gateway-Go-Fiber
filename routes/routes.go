package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seccret404/BE-payment-gateway-Go-Fiber/handlers"
	"gorm.io/gorm"
)

func AppRoute(app *fiber.App, db *gorm.DB){
	route := app.Group("/api")

	route.Post("/payment", handlers.CreatePayment(db))
}