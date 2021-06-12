package app

import (
	"fiber-boilerplate/pkg/global"
	"fiber-boilerplate/pkg/initial"

	"github.com/gofiber/fiber/v2"
)

func Bootstrap() {
	app := fiber.New(fiber.Config{
		ErrorHandler: catchUnknownError,
	})

	initial.InitApp(app)

	app.Listen(":" + global.ServerSetting.HttpPort)
}

func catchUnknownError(ctx *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ctx.Status(code).SendString("Uncatch Error")
}
