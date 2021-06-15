package middleware

import (
	"fmt"
	"runtime"

	"fiber-boilerplate/pkg/errcode"
	"fiber-boilerplate/pkg/response"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func NewRecovery() fiber.Handler {
	return func(ctx *fiber.Ctx) (err error) {
		defer func() {
			if e := recover(); e != nil {
				buf := make([]byte, 2048)
				buf = buf[:runtime.Stack(buf, false)]

				zap.S().Errorf("%v>>%s\n", e, buf)

				// 判断是否是内部已定义错误
				if _, ok := e.(*errcode.Error); ok {
					response.New(ctx).ToErrorResponse(e.(*errcode.Error))
				} else {
					// 不是已定义错误，交给统一错误处理
					err = fmt.Errorf("%v", e)
				}
			}
		}()

		return ctx.Next()
	}
}
