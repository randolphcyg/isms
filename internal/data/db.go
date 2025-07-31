package data

import (
	"fmt"
	"isms/internal/conf"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(cfg *conf.Data_Database) *gorm.DB {
	// 构建MySQL连接字符串
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 应用连接池配置
	sqlDB.SetMaxOpenConns(int(cfg.MaxOpenConns))
	sqlDB.SetMaxIdleConns(int(cfg.MaxIdleConns))
	sqlDB.SetConnMaxLifetime(cfg.ConnMaxLifetime.AsDuration())
	sqlDB.SetConnMaxIdleTime(cfg.ConnMaxIdleTime.AsDuration())

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	return db
}
