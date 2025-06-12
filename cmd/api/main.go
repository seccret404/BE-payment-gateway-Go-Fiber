package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/seccret404/BE-payment-gateway-Go-Fiber/config"
)

func main() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Cannot load file .env")
	}
	config.ConnectDB()

	app := fiber.New()

	log.Fatal(app.Listen(":3000"))
}