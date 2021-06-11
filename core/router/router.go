package router

import (
	"github.com/gofiber/fiber/v2"

	v1 "fiber-boilerplate/core/handler/v1"
)

func NewRouter(app *fiber.App) {
	tag := v1.NewTag()
	article := v1.NewArticle()

	apiv1 := app.Group("/api/v1")
	{
		apiv1.Post("/tags", tag.Create)
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