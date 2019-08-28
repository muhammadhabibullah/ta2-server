package models

import (
	"tugas-akhir-2/common"

	"github.com/jinzhu/gorm"
)

//Target model
type Target struct {
	gorm.Model
	Name         string  `gorm:"column:name;not null"`
	TargetType   string  `gorm:"column:targettype;type:enum('T','D','K','P','E','');not null"`
	TargetNumber float64 `gorm:"column:targetnumber;type:float;not null"`
	UserID       uint    `gorm:"column:userid;not null"`
}

// Serialize serializes bicycle data
func (t *Target) Serialize() common.JSON {
	return common.JSON{
		"id":           t.ID,
		"name":         t.Name,
		"targettype":   t.TargetType,
		"targetnumber": t.TargetNumber,
		"userid":       t.UserID,
	}
}

func (t *Target) Read(m common.JSON) {
	t.ID = uint(m["id"].(float64))
	t.Name = m["name"].(string)
	t.TargetType = m["targettype"].(string)
	t.TargetNumber = m["targetnumber"].(float64)
	t.UserID = uint(m["userid"].(float64))
}
