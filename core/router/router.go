package router

import (
	"github.com/gofiber/fiber/v2"

	v1 "fiber-boilerplate/core/handler/v1"
)

func BindRouter(app *fiber.App) {
	// /api/v1
	apiv1Group := app.Group("/api/v1")

	// /api/v1/car
	carGroup := apiv1Group.Group("/car")
	{
		car := v1.NewCar()

		carGroup.Get("/", car.List)
	}
}
