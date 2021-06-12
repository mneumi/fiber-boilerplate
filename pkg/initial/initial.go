package initial

import (
	"time"

	"fiber-boilerplate/core/dao"
	"fiber-boilerplate/core/model"
	"fiber-boilerplate/core/router"
	"fiber-boilerplate/pkg/global"
	"fiber-boilerplate/pkg/logger"
	"fiber-boilerplate/pkg/setting"

	"github.com/gofiber/fiber/v2"
)

func InitApp(app *fiber.App) {
	global.App = app

	initSetting()
	initRouter()
	initServer()
	initMiddleware()
	initLogger()
	initDB()
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

func initMiddleware() {}
