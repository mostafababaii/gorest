package models

import (
	"fmt"
	"log"

	"github.com/mostafababaii/gorest/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Mapper interface {
	MapTo(p any)
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
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	defer Migrate()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("issue on getting db instance err: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
}

func Migrate() {
	db.AutoMigrate(&User{})
}
