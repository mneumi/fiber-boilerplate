package initial

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"fiber-boilerplate/core/dao"
	"fiber-boilerplate/core/model"
	"fiber-boilerplate/core/router"
	"fiber-boilerplate/pkg/global"
	"fiber-boilerplate/pkg/logger"
	"fiber-boilerplate/pkg/middleware"
	"fiber-boilerplate/pkg/setting"

	"github.com/bwmarrin/snowflake"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func InitApp(app *fiber.App) {
	global.App = app

	initSetting()
	initGracefulShutdown()
	initServer()
	initValidator()
	initLogger()
	initDB()
	initSnowflakeNode()
	initMiddleware()
	initRouter()
}

func initGracefulShutdown() {
	// 新建管道
	ch := make(chan os.Signal, 1)
	// 注册事件，关闭程序时，向管道发送一个信号
	signal.Notify(ch, os.Interrupt)

	go func() {
		// 阻塞等待信号
		<-ch
		fmt.Println("Gracefully shutdown ...")
		_ = global.App.Shutdown()
	}()
}

func initSetting() {
	setting, err := setting.NewSetting()
	if err != nil {
		panic(err)
	}

	err = setting.ReadSection("server", &global.ServerSetting)
	if err != nil {
		panic(err)
	}

	err = setting.ReadSection("app", &global.AppSetting)
	if err != nil {
		panic(err)
	}

	err = setting.ReadSection("database", &global.DatabaseSetting)
	if err != nil {
		panic(err)
	}

	err = setting.ReadSection("logger", &global.LoggerSetting)
	if err != nil {
		panic(err)
	}

	err = setting.ReadSection("routerLimiter", &global.RouterLimiter)
	if err != nil {
		panic(err)
	}

	err = setting.ReadSection("jwt", &global.JWTSetting)
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

func initSnowflakeNode() {
	var err error
	global.SnowflakeNode, err = snowflake.NewNode(1)

	if err != nil {
		panic(err)
	}
}

func initMiddleware() {
	global.App.Use(middleware.NewRecovery())
	global.App.Use(middleware.NewAttackLimiter())
	global.App.Use(middleware.NewRouterLimiter())
	global.App.Use(middleware.NewCors())
	global.App.Use(middleware.NewAppInfo())
	global.App.Use(middleware.NewRequestID())
	global.App.Use(middleware.NewAccessLogger())
}
