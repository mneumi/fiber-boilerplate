package middleware

import (
	"strings"

	"fiber-boilerplate/pkg/errcode"
	"fiber-boilerplate/pkg/global"
	"fiber-boilerplate/pkg/response"

	"github.com/gofiber/fiber/v2"
)

func Validate(dto interface{}) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if err := ctx.BodyParser(dto); err != nil {
			return response.New(ctx).
				ToErrorResponse(errcode.InvalidParams.
					WithDetails(err.Error()))
		}

		if err := global.Validator.Struct(dto); err != nil {
			errs := strings.Split(err.Error(), "\n")

			return response.New(ctx).
				ToErrorResponse(errcode.InvalidParams.
					WithDetails(errs...))
		}

		return ctx.Next()
	}
}
