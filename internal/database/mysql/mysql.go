package mysql

import (
	"fmt"
	"log"

	"github.com/mostafababaii/gorest/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func NewConnection() *gorm.DB {
	return Connection
}

// Setup initializes the database instance
func Setup() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseConfig.User,
		config.DatabaseConfig.Password,
		config.DatabaseConfig.Host,
		config.DatabaseConfig.Name,
	)

	var err error
	Connection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	sqlDB, err := Connection.DB()
	if err != nil {
		log.Fatalf("issue on getting db instance err: %v", err)
	}

	sqlDB.SetMaxIdleConns(config.DatabaseConfig.MaxIdle)
	sqlDB.SetMaxOpenConns(config.DatabaseConfig.MaxOpen)
}
