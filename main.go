package main

import (
	"log"
	"time"

	"fiber-boilerplate/core/router"
	"fiber-boilerplate/global"
	"fiber-boilerplate/pkg/setting"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Server().ReadTimeout = global.ServerSetting.ReadTimeout
	app.Server().WriteTimeout = global.ServerSetting.WriteTimeout

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World !")
	})

	router.NewRouter(app)

	app.Listen(":" + global.ServerSetting.HttpPort)
}

func init() {
	err := setupSetting()

	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting()

	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)

	if err != nil {
		return err
	}

	err = setting.ReadSection("App", &global.AppSetting)

	if err != nil {
		return err
	}

	err = setting.ReadSection("Database", &global.DatabaseSetting)

	if err != nil {
		return err
	}

	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}
