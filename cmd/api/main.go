package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/seccret404/BE-payment-gateway-Go-Fiber/config"
	// "github.com/seccret404/BE-payment-gateway-Go-Fiber/models"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Cannot load file .env")
	}
	config.ConnectDB()
	// models.Migrate()

	app := fiber.New()

	log.Fatal(app.Listen(":3000"))
}