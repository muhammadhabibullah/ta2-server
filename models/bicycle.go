package models

import (
	"fmt"
	"tugas-akhir-2/common"
)

//Bicycle model
type Bicycle struct {
	ID       uint   `gorm:"column:id;primary_key"`
	Name     string `gorm:"column:name;not null"`
	BikeType string `gorm:"column:biketype;not null"`
	UserID   uint   `gorm:"column:userid;not null"`
}

// Serialize serializes bicycle data
func (b *Bicycle) Serialize() common.JSON {
	return common.JSON{
		"id":       fmt.Sprint(b.ID),
		"name":     b.Name,
		"biketype": b.BikeType,
		"userid":   fmt.Sprint(b.UserID),
	}
}

func (b *Bicycle) Read(m common.JSON) {
	b.ID = uint(m["id"].(float64))
	b.Name = m["name"].(string)
	b.BikeType = m["biketype"].(string)
	b.UserID = uint(m["userid"].(float64))
}
