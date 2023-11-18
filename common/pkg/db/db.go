package db

import (
	"fmt"

	"github.com/nilsyadv/ShopBillBuddy/common/pkg/config"
	"github.com/nilsyadv/ShopBillBuddy/common/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config config.InterfaceConfig, logger logger.InterfaceLogger) (*gorm.DB, error) {
	switch config.GetString("db.type") {
	case "Mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			config.GetString("user"), config.GetString("pass"), config.GetString("host"), config.GetString("port"), config.GetString("dbname"))
		if ssl := config.GetString("db.sslmode"); ssl != "" {
			dsn += " sslmode=" + ssl
		}
		if tz := config.GetString("db.timezone"); tz != "" {
			dsn += " TimeZone" + tz
		}
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Error("encounter during db connection", err)
			return db, err
		}
		return db, nil
	case "Postgresql":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
			config.GetString("host"), config.GetString("user"), config.GetString("pass"), config.GetString("dbname"), config.GetString("port"))
		if ssl := config.GetString("db.sslmode"); ssl != "" {
			dsn += " sslmode=" + ssl
		}
		if tz := config.GetString("db.timezone"); tz != "" {
			dsn += " TimeZone" + tz
		}
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Error("encounter during db connection", err)
			return db, err
		}
		return db, nil
	default:
		return nil, fmt.Errorf("unknown db type %s", config.GetString("db.type"))
	}
}
