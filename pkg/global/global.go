package global

import (
	"time"

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
)

type ServerSettingSection struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingSection struct {
	DefaultPageSize int
	MaxPageSize     int
}

type DatabaseSettingSection struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	Port         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    string
	Loc          string
	MaxIdleConns int
	MaxOpenConns int
}

type LoggerSettingSection struct {
	FilePath   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	LocalTime  bool
}
