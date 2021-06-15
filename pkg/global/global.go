package global

import (
	"github.com/bwmarrin/snowflake"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/juju/ratelimit"
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

	// Snowflake Node
	SnowflakeNode *snowflake.Node

	// Config String
	DevConfigYAMLContent string

	// RouterGroupLimiter
	RouterLimiterMap = make(map[string]*ratelimit.Bucket)
	RouterLimiter    *RouterLimiterSection

	// JWT
	JWTSetting *JWTSettingSection
)
