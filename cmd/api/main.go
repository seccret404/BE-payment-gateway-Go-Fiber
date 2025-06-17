package main

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"app/config"
	"app/routes"
	"app/models"
)


func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Cannot load file .env")
	}
	config.ConnectDB()
	models.Migrate()

	app := fiber.New()
	//cors
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", //fe url
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	
	routes.AppRoute(app, config.DB)

	

	app.Listen("0.0.0.0:3000")

}
