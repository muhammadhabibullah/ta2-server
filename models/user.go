package models

import (
	"time"

	"tugas-akhir-2/common"

	"github.com/jinzhu/gorm"
)

//User model
type User struct {
	gorm.Model
	Email          string     `gorm:"column:email;unique;not null"`
	HashedPassword string     `gorm:"column:password;not null"`
	Name           string     `gorm:"column:name;not null"`
	Birthdate      *time.Time `gorm:"birthdate;type:date;not null"`
	Gender         string     `gorm:"gender;not null"`
	Weight         float32    `gorm:"column:weight;type:float;not null"`
	Height         float32    `gorm:"column:height;type:float;not null"`
}

// Serialize serializes user data
func (u *User) Serialize() common.JSON {
	return common.JSON{
		"id":        u.ID,
		"email":     u.Email,
		"name":      u.Name,
		"birthdate": u.Birthdate,
		"gender":    u.Gender,
		"weight":    u.Weight,
		"height":    u.Height,
	}
}

func (u *User) Read(m common.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Name = m["name"].(string)
	u.Birthdate = m["birthdate"].(*time.Time)
	u.Gender = m["gender"].(string)
	u.Weight = m["weight"].(float32)
	u.Height = m["height"].(float32)
}
