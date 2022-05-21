package storage

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DataBase ...
type DataBase struct {
	DriverName string
	Host       string
	Port       int
	User       string
	PassWord   string
	DBName     string
}

// RedisConf ...
type RedisConf struct {
	NetWork     string
	HostAddr    string
	PassWord    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout int64
}

// NewGormDB new gorm.DB
func NewGormDB(c *DataBase) *gorm.DB {
	// select the database type as mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.User, c.PassWord, c.Host, c.Port, c.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(10)
	return db
}

// RecordNotFound ...
func RecordNotFound(err error) bool {
	return err == gorm.ErrRecordNotFound
}
