package controllers

import (
	"tugas-akhir-2/common"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//GetData function
func GetData(c *gin.Context) {

	var bicycle Bicycle
	var user User
	var targetd Target
	var targete Target
	var targett Target

	db := c.MustGet("db").(*gorm.DB)

	bicycleid := c.Param("bid")
	db.Where("id = ?", bicycleid).Find(&bicycle)

	userid := bicycle.UserID
	db.Where("id = ?", userid).Find(&user)

	db.Where("userid = ? and targettype = 'D'", userid).Find(&targetd).Order("created_by DESC").Limit(1)
	db.Where("userid = ? and targettype = 'E'", userid).Find(&targete).Order("created_by DESC").Limit(1)
	db.Where("userid = ? and targettype = 'T'", userid).Find(&targett).Order("created_by DESC").Limit(1)

	age := user.CountAge()

	c.JSON(200, common.JSON{
		"age":       age,
		"distance":  targetd.TargetNumber,
		"elevation": targete.TargetNumber,
		"time":      targett.TargetNumber,
	})

}
