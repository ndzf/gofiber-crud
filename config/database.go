package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var databaseConfig map[string]interface{}

func Connect() {
	var err error

	databaseConfig = map[string]interface{}{
		"DB_HOST":     "localhost",
		"DB_NAME":     "gofiber_crud",
		"DB_USER":     "postgres",
		"DB_PASSWORD": "postgres",
		"DB_PORT":     5432,
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", databaseConfig["DB_HOST"], databaseConfig["DB_PORT"], databaseConfig["DB_USER"], databaseConfig["DB_PASSWORD"], databaseConfig["DB_NAME"])
	DB, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed connect to database")
	}

	fmt.Println(DB)
}
