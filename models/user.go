package models

import (
	"bytes"
	"fmt"
	"time"
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

//CountAge return age of the user
func (u *User) CountAge() int {

	var b bytes.Buffer

	bd := u.Birthdate[:len(u.Birthdate)-6] //cut +07:00 char
	b.WriteString(bd)
	b.WriteString(".000Z") // add .000Z char
	bd = b.String()

	t, err := time.Parse("2006-01-02T15:04:05.000Z", bd)
	if err != nil {
		fmt.Println(err)
	}

	y1, m1, d1 := t.Date()
	y2, m2, d2 := time.Now().Date()
	year := int(y2 - y1)
	month := int(m2 - m1)
	day := int(d2 - d1)

	if day < 0 {
		// days in month:
		t := time.Date(y1, m1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}
	return year
}
