package middleware

import "github.com/gofiber/fiber/v2"

func NewAppInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("version", "1.0.0")
		ctx.Locals("app_name", "fiber-boilerplate")

		return ctx.Next()
	}
}
