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
type Target = models.Target

//TargetSignUp to input your target
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
		"target": target.Serialize(),
	})

}

//LastestTargetRetrieve return lastest target by user
func LastestTargetRetrieve(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	var target Target

	db := c.MustGet("db").(*gorm.DB)
	//db.Raw("SELECT * FROM targets WHERE userid = ? ORDER BY id DESC LIMIT 1", user.ID).Scan(&target)
	db.Limit(1).Where("userid = ?", user.ID).Order("id desc").Find(&target)

	c.JSON(200, target.Serialize())
}

//TargetEdit edit target detail
func TargetEdit(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	type RequestBody struct {
		Name         string  `json:"name"`
		TargetType   string  `json:"targettype" binding:"required"`
		TargetNumber float64 `json:"targetnumber" binding:"required"`
	}

	var body RequestBody

	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	var target Target

	if err := db.Where("id = ?", id).First(&target).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if target.UserID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	target.Name = body.Name
	target.TargetType = body.TargetType
	target.TargetNumber = body.TargetNumber
	db.Save(&target)
	c.JSON(200, target.Serialize())
}

//TargetDelete delete target
func TargetDelete(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")

	var target Target

	if err := db.Where("id = ?", id).First(&target).Error; err != nil {
		c.AbortWithStatus(404)
		return
	}

	if target.UserID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	db.Delete(&target)
	c.Status(204)
}
