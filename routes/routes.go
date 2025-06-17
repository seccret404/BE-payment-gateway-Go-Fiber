package routes

import (
	"github.com/gofiber/fiber/v2"
	"app/handlers"
	"gorm.io/gorm"
)

func AppRoute(app *fiber.App, db *gorm.DB){
	route := app.Group("/api")

	route.Post("/payment", handlers.CreatePayment(db))
}