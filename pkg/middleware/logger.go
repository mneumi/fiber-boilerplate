package middleware

import (
	"fiber-boilerplate/pkg/global"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewAccessLogger() fiber.Handler {
	formatString := `{"date":"${time}","method":"${method}","path":"${path}","status":"${status}","latency":"${latency}","requestid":"${locals:requestid}","ip":"${ip}"}` + "\n"

	return logger.New(logger.Config{
		Format:     formatString,
		TimeFormat: "2006-01-02",
		TimeZone:   "Local",
		Output: &lumberjack.Logger{
			Filename:   global.LoggerSetting.CommonPath,
			MaxSize:    global.LoggerSetting.MaxSize,
			MaxBackups: global.LoggerSetting.MaxBackups,
			MaxAge:     global.LoggerSetting.MaxAge,
			Compress:   global.LoggerSetting.Compress,
		},
	})
}
