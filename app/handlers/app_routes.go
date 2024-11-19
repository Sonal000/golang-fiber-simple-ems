package handler

import (
	middleware "employeeManagement/app/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	users := api.Group("/users", middleware.AuthProtect())
	RegisterUserRoutes(users)

	emp := api.Group("/emp", middleware.AuthProtect())
	RegisterEmpRoutes(emp)

	auth := api.Group("/auth")
	RegisterAuthRoutes(auth)
}

func RegisterAuthRoutes(auth fiber.Router) {
	authHandler := NewAuthHandler()

	auth.Post("/login", authHandler.Login)
	auth.Post("/register", authHandler.RegisterAdmin)
}

func RegisterUserRoutes(users fiber.Router) {
	userHandler := NewUserHandler()

	users.Post("/", middleware.AuthRestrict("admin"), userHandler.CreateUser)
	users.Get("/", middleware.AuthRestrict("admin"), userHandler.GetUsers)
	users.Get("/:id", middleware.AuthRestrict("admin"), userHandler.GetUser)
	users.Put("/:id", middleware.AuthRestrict("admin"), userHandler.UpdateUser)
}

func RegisterEmpRoutes(emp fiber.Router) {
	empHandler := NewEmployeeHandler()
	emp.Get("/profile", middleware.AuthRestrict("user"), empHandler.GetEmployeeProfile)
	emp.Put("/profile", middleware.AuthRestrict("user"), empHandler.UpdateEmployeeProfile)
	emp.Post("/", middleware.AuthRestrict("admin"), empHandler.CreateEmployee)
	emp.Get("/", middleware.AuthRestrict("admin"), empHandler.GetEmployees)
	emp.Get("/:id", middleware.AuthRestrict("admin"), empHandler.GetEmployee)
	emp.Put("/:id", middleware.AuthRestrict("admin"), empHandler.UpdateEmployee)
	emp.Delete("/:id", middleware.AuthRestrict("admin"), empHandler.DeleteEmployee)
}
