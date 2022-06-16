package lib

import (
	"github.com/m0hammedimran/todo_server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func db() *gorm.DB {
	dsn := "host=localhost user=user password=password dbname=todos port=5432 sslmode=disable"
	postgres, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	postgres.AutoMigrate(&models.Todo{})
	return postgres
}

var DB = db()
