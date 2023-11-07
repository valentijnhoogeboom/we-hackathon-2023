package database

import (
	"GlobalAPI/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var StaticDatabase *Database

type Database struct {
	Host     string   `json:"host"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Name     string   `json:"database"`
	DB       *gorm.DB `json:"db"`
}

func (database *Database) Connect() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", database.Username, database.Password, database.Host, database.Name)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	database.DB = db
}

func (database *Database) Migrate() {
	if database.DB == nil {
		database.Connect()
	}

	err := database.DB.AutoMigrate(&models.User{}, &models.Meter{}, &models.Data{})
	if err != nil {
		panic(err)
		return
	}
}
