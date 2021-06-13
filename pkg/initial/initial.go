package initial

import (
	"fmt"
	"runtime"
	"time"

	"fiber-boilerplate/core/dao"
	"fiber-boilerplate/core/model"
	"fiber-boilerplate/core/router"
	"fiber-boilerplate/pkg/errcode"
	"fiber-boilerplate/pkg/global"
	"fiber-boilerplate/pkg/logger"
	"fiber-boilerplate/pkg/response"
	"fiber-boilerplate/pkg/setting"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

func InitApp(app *fiber.App) {
	global.App = app

	initSetting()
	initServer()
	initValidator()
	initLogger()
	initDB()
	initMiddleware()
	initRouter()
}

func initSetting() {
	setting, err := setting.NewSetting()
	if err != nil {
		panic(err)
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		panic(err)
	}

	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		panic(err)
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		panic(err)
	}

	err = setting.ReadSection("Logger", &global.LoggerSetting)
	if err != nil {
		panic(err)
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
}

func initRouter() {
	router.BindRouter(global.App)
}

func initServer() {
	global.App.Server().ReadTimeout = global.ServerSetting.ReadTimeout
	global.App.Server().WriteTimeout = global.ServerSetting.WriteTimeout
}

func initValidator() {
	global.Validator = validator.New()
}

func initDB() {
	db, err := dao.NewDB(global.DatabaseSetting)

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Article{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.ArticleTag{})

	global.DB = db
}

func initLogger() {
	logger.InitLogger(global.LoggerSetting)
}

func initMiddleware() {
	global.App.Use(newRecoveryMiddleware())
	global.App.Use(newLimiterMiddleware())
	global.App.Use(newAppInfo())
	global.App.Use(cors.New())
	global.App.Use(requestid.New())
	global.App.Use(logger.New())
}

func newLimiterMiddleware() fiber.Handler {
	return limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        30,
		Expiration: 30 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusTooManyRequests)
		},
	})
}

func newRecoveryMiddleware() fiber.Handler {
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

func newAppInfo() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Locals("version", "1.0.0")
		ctx.Locals("app_name", "fiber-boilerplate")

		return ctx.Next()
	}
}
