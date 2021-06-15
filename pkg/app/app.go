package app

import (
	"fiber-boilerplate/pkg/global"
	"fiber-boilerplate/pkg/initial"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
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
	} else {
		zap.S().Errorf("Uncatch Error: %v", err.Error())
	}

	return ctx.Status(code).JSON(fiber.Map{
		"msg": fmt.Sprintf("Uncatch Error: %s", err.Error()),
	})
}
