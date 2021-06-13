package logger

import (
	"fiber-boilerplate/pkg/global"

	"github.com/gofiber/fiber/v2"
	loggerMiddleware "github.com/gofiber/fiber/v2/middleware/logger"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 这个 New 是新建中间件，该中间件记录每一次的的请求和响应，属于自动记录
func New() func(*fiber.Ctx) error {
	formatString := `{"date":"${time}","method":"${method}","path":"${path}","status":"${status}","latency":"${latency}","requestid":"${locals:requestid}","ip":"${ip}"}` + "\n"

	return loggerMiddleware.New(loggerMiddleware.Config{
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

func InitLogger(loggerSetting *global.LoggerSettingSection) {
	encoder := getEncoder()

	writeSyncer := getLogWriter(
		loggerSetting.FilePath,
		loggerSetting.MaxSize,
		loggerSetting.MaxBackups,
		loggerSetting.MaxAge,
		loggerSetting.Compress,
	)

	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filePath string, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filePath,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}

	return zapcore.AddSync(lumberJackLogger)
}
