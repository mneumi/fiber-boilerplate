package global

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	// App
	App *fiber.App

	// Setting
	ServerSetting   *ServerSettingSection
	AppSetting      *AppSettingSection
	DatabaseSetting *DatabaseSettingSection
	LoggerSetting   *LoggerSettingSection

	// DB
	DB *gorm.DB

	// Validator
	Validator *validator.Validate
)
