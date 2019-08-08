package models

import (
	"tugas-akhir-2/common"

	"github.com/jinzhu/gorm"
)

//User model
type User struct {
	gorm.Model
	Email          string  `gorm:"column:email;unique;not null"`
	HashedPassword string  `gorm:"column:password;not null"`
	Name           string  `gorm:"column:name;not null"`
	Birthdate      string  `gorm:"column:birthdate;type:date;not null"`
	Gender         string  `gorm:"column:gender;type:enum('M','F','');not null"`
	Weight         float64 `gorm:"column:weight;type:float;not null"`
	Height         float64 `gorm:"column:height;type:float;not null"`
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
	u.Email = m["email"].(string)
	u.Name = m["name"].(string)
	u.Birthdate = m["birthdate"].(string)
	u.Gender = m["gender"].(string)
	u.Weight = m["weight"].(float64)
	u.Height = m["height"].(float64)
}
