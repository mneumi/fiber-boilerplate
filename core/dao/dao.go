package dao

import (
	"fmt"

	"fiber-boilerplate/pkg/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(dbSetting *global.DatabaseSettingSection) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		global.DatabaseSetting.UserName,
		global.DatabaseSetting.Password, global.DatabaseSetting.Host,
		global.DatabaseSetting.Port,
		global.DatabaseSetting.DBName,
		global.DatabaseSetting.Charset,
		global.DatabaseSetting.ParseTime,
		global.DatabaseSetting.Loc,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, err
	}

	maxIdleConns := global.DatabaseSetting.MaxIdleConns
	maxOpenConns := global.DatabaseSetting.MaxOpenConns

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)

	return db, nil
}
