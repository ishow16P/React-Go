package routes

import (
	"go-api/controller"
	"go-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	// if under middleware will use middleware
	app.Use(middleware.IsAuthenticate)
}
