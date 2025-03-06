package database

import (
	"fmt"

	"tartalom/config"
	"tartalom/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() {
	var err error

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.Config("DATABASE_USER"),
		config.Config("DATABASE_PASS"),
		config.Config("DATABASE_HOST"),
		config.Config("DATABASE_PORT"),
		config.Config("DATABASE_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connected successfully!")
	db.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
	DB = db
}
