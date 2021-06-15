package middleware

import (
	"fiber-boilerplate/pkg/global"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func NewRequestID() fiber.Handler {
	return requestid.New(requestid.Config{
		Generator: func() string {
			return "r" + global.SnowflakeNode.Generate().String()
		},
	})
}
