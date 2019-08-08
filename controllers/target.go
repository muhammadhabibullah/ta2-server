package controllers

import (
	"fmt"
	"tugas-akhir-2/common"
	"tugas-akhir-2/middlewares"
	"tugas-akhir-2/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Target is alias for models.Target
type Target models.Target

func TargetSignUp(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		Name         string  `json:"name" binding:"required"`
		TargetType   string  `json:"targettype" binding:"required"`
		TargetNumber float64 `json:"targetnumber" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	target := Target{
		Name:         body.Name,
		TargetType:   body.TargetType,
		TargetNumber: body.TargetNumber,
		UserID:       user.ID,
	}

	db.NewRecord(target)
	db.Create(&target)

	c.JSON(200, common.JSON{
		"target": target,
	})

}

func LastestTargetRetrieve(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	var target Target

	db := c.MustGet("db").(*gorm.DB)
	db.Raw("SELECT * FROM targets WHERE userid = ? ORDER BY id DESC LIMIT 1", user.ID).Scan(&target)

	c.JSON(200, target)
}
