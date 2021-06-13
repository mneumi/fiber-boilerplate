package router

import (
	"github.com/gofiber/fiber/v2"

	"fiber-boilerplate/core/dto"
	v1 "fiber-boilerplate/core/handler/v1"
	"fiber-boilerplate/core/middleware"
)

func BindRouter(app *fiber.App) {
	tag := v1.NewTag()
	article := v1.NewArticle()

	apiv1 := app.Group("/api/v1")
	{
		apiv1.Get("/", func(ctx *fiber.Ctx) error {
			panic("I'm an error")
		})

		apiv1.Post("/tags", middleware.Validate(&dto.UserDto{}), tag.Create)
		apiv1.Delete("/tags/:id", tag.Delete)
		apiv1.Put("/tags/:id", tag.Update)
		apiv1.Patch("/tags/:id/state", tag.Update)
		apiv1.Get("/tags", tag.List)

		apiv1.Post("/articles", article.Create)
		apiv1.Delete("/articles/:id", article.Delete)
		apiv1.Put("/articles/:id", article.Update)
		apiv1.Patch("/articles/:id/state", article.Update)
		apiv1.Get("/articles", article.List)
	}
}
