// main.go
package main

import (
	"CRUD-Fiber/database"
	"CRUD-Fiber/handlers"
	"CRUD-Fiber/models"
	"CRUD-Fiber/repositories"
	"CRUD-Fiber/services"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Connect to the database
	db, err := database.Connect()
	if err != nil {
		panic("Failed to connect to the database")
	}

	// AutoMigrate to create all tables in the models package
	modelsToMigrate := []interface{}{
		&models.User{},
		// Add other model structs here
	}

	for _, model := range modelsToMigrate {
		if err := db.AutoMigrate(model); err != nil {
			panic("Failed to auto-migrate the database")
		}
	}
	defer func() {
		// Close the database connection when the application exits
		sqlDB, err := db.DB()
		if err != nil {
			panic("Failed to get underlying DB instance")
		}
		sqlDB.Close()
	}()

	// Initialize the repository and service
	userRepository := repositories.NewGormUserRepository(db)
	userService := services.NewUserServiceImpl(userRepository)

	// Initialize the handlers
	userHandler := handlers.NewUserHandler(userService)

	routesAPI(app, userHandler)

	// Start the server
	app.Listen(":3000")
}

func routesAPI(app *fiber.App, userHandler *handlers.UserHandler) {
	// API Routes
	app.Get("/users", userHandler.GetAllUsers)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Post("/users", userHandler.CreateUser)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
}
