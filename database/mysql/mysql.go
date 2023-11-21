package mysql

import (
	"fmt"

	"github.com/fazaalexander/montirku-be/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	InitDB()
	// InitialMigration()
}

func InitDB() {
	var err error

	config := config.GetConfig()

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println(DB)
}

func InitialMigration() {

}
