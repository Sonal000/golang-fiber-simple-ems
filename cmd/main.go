package main

import (
	handler "employeeManagement/app/handlers"
	db "employeeManagement/pkg"
	"log"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	projectRoot, err := filepath.Abs(filepath.Join(".", ".."))
	if err != nil {
		log.Fatalf("Error finding project root: %v", err)
	}

	envFilePath := filepath.Join(projectRoot, ".env")
	err = godotenv.Load(envFilePath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	err = db.InitDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	start()
}

func start() {
	app := fiber.New()

	handler.RegisterRoutes(app)

	err := app.Listen(":8181")
	if err != nil {
		log.Fatal(err)
	}
}
