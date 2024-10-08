package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// Migrate automigrates models using ORM
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Bicycle{})
	db.AutoMigrate(&Target{})
	db.AutoMigrate(&Cycling{})
	db.Model(&Bicycle{}).AddForeignKey("userid", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&Target{}).AddForeignKey("userid", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&Cycling{}).AddForeignKey("userid", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&Cycling{}).AddForeignKey("bicycleid", "bicycles(id)", "RESTRICT", "RESTRICT")
	// set up foreign keys
	fmt.Println("Auto Migration has beed processed")
}
