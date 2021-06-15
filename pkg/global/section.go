package global

import "time"

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
	CommonPath string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	LocalTime  bool
}

type RouterLimiterSection struct {
	Rules []RouterLimiterItem
}

type RouterLimiterItem struct {
	Path     string
	Interval int64
	Capacity int64
	Quantum  int64
}

type JWTSettingSection struct {
	Secret string
	Issuer string
	Expire int
}
