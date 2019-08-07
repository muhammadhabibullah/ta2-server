package database

import (
	"fmt"
	"os"

	"tugas-akhir-2/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql" //
)

// CreateConnection DB
func CreateConnection() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_CONFIG"))
	db.LogMode(true) // logs SQL
	if err != nil {
		fmt.Print(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	models.Migrate(db)
	return db
}

// Inject database to gin context
func Inject(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}
