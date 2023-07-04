package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/joshjms/auth-service/middlewares"
	"github.com/joshjms/auth-service/routes"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}

	app := fiber.New()

	app.Post("/register", routes.Register)
	app.Post("/login", routes.Login)
	app.Post("/refresh", middlewares.AuthMiddleware(os.Getenv("REFRESH_TOKEN_SIGNATURE")), routes.Refresh)
	app.Get("/verify", middlewares.AuthMiddleware(os.Getenv("ACCESS_TOKEN_SIGNATURE")), routes.ProtectedRoute)

	log.Fatal(app.Listen(":3001"))
}
