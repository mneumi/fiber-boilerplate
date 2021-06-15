package middleware

import (
	"fiber-boilerplate/pkg/errcode"
	"fiber-boilerplate/pkg/response"
	jwtutil "fiber-boilerplate/pkg/utils/jwt"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func JWT() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var token string = ctx.Get("token")
		var ecode *errcode.Error

		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			claims, err := jwtutil.ParseToken(token)

			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			} else {
				ctx.Set("username", claims.Username)
			}
		}

		if ecode != nil {
			return response.New(ctx).ToErrorResponse(ecode)
		}

		return ctx.Next()
	}
}
