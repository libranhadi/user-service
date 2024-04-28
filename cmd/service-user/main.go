package main

import (
	"fmt"
	"service-user/controller"
	"service-user/database"
	"service-user/middleware"
	"service-user/repository"
	"service-user/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.InitDatabase()
	userRepo := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)
	middleware := middleware.NewAuthImpl(userRepo)

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return nil
	})

	users := app.Group("/users")

	users.Post("/register", userController.Register)
	users.Post("/login", userController.Login)
	users.Get("/auth", middleware.Authentication, userController.Auth)

	port := 3000
	fmt.Printf("Service user is running on port:%d...\n ", port)
	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Printf("Error starting Service user: %v\n", err)
	}
}
