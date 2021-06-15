package middleware

import (
	"fiber-boilerplate/pkg/errcode"
	"fiber-boilerplate/pkg/global"
	"fiber-boilerplate/pkg/response"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/juju/ratelimit"
)

func NewRouterLimiter() fiber.Handler {
	prefixArray := []string{}

	for _, item := range global.RouterLimiter.Rules {
		prefixArray = append(prefixArray, item.Path)
		global.RouterLimiterMap[item.Path] = ratelimit.NewBucketWithQuantum(time.Duration(item.Interval)*time.Second, item.Capacity, item.Quantum)
	}

	return func(ctx *fiber.Ctx) error {
		uri := string(ctx.Request().URI().Path())

		for _, prefix := range prefixArray {
			if strings.Contains(uri, prefix) {

				bucket := global.RouterLimiterMap[prefix]

				count := bucket.TakeAvailable(1)

				if count == 0 {
					return response.New(ctx).ToErrorResponse(errcode.TooManyRequests)
				}

				break
			}
		}

		return ctx.Next()
	}
}
