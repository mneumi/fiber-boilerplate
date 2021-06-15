package router

import (
	"github.com/gofiber/fiber/v2"

	"fiber-boilerplate/core/dto"
	v1 "fiber-boilerplate/core/handler/v1"
	"fiber-boilerplate/pkg/middleware"
)

func BindRouter(app *fiber.App) {
	tag := v1.NewTag()
	article := v1.NewArticle()

	// /api/v1
	apiv1Group := app.Group("/api/v1")
	{
		apiv1Group.Get("/", func(ctx *fiber.Ctx) error {
			panic("I'm an error")
		})
	}

	// /api/v1/tags
	tagsGroup := apiv1Group.Group("/tags")
	{
		tagsGroup.Get("/", tag.List)
		tagsGroup.Post("/", middleware.Validate(&dto.UserDto{}), tag.Create)
		tagsGroup.Delete("/:id", tag.Delete)
		tagsGroup.Put("/:id", tag.Update)
		tagsGroup.Patch("/:id/state", tag.Update)
	}

	// /api/v1/articles
	apiv1Group.Post("/articles", article.Create)
	apiv1Group.Delete("/articles/:id", article.Delete)
	apiv1Group.Put("/articles/:id", article.Update)
	apiv1Group.Patch("/articles/:id/state", article.Update)
	apiv1Group.Get("/articles", article.List)
}
