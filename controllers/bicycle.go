package controllers

import (
	"fmt"
	"tugas-akhir-2/common"
	"tugas-akhir-2/middlewares"
	"tugas-akhir-2/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Bicycle is alias for models.Bicycle
type Bicycle = models.Bicycle

//BicycleSignUp to input your bicycle
func BicycleSignUp(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	db := c.MustGet("db").(*gorm.DB)
	type RequestBody struct {
		ID       uint   `json:"id" binding:"required"`
		Name     string `json:"name"`
		BikeType string `json:"biketype" binding:"required"`
	}

	var body RequestBody
	if err := c.BindJSON(&body); err != nil {
		fmt.Println(err)
		c.AbortWithStatus(400)
		return
	}

	// check existancy
	var exists Bicycle
	if err := db.Where("ID = ?", body.ID).First(&exists).Error; err == nil {
		c.AbortWithStatus(409)
		return
	}

	//create bicycle
	bicycle := Bicycle{
		ID:       body.ID,
		Name:     body.Name,
		BikeType: body.BikeType,
		UserID:   user.ID,
	}

	db.NewRecord(bicycle)
	db.Create(&bicycle)

	c.JSON(200, common.JSON{
		"bicycle": bicycle.Serialize(),
	})
}

//BicycleRetrieve get all bicycles user
func BicycleRetrieve(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	var bicycles []Bicycle

	db := c.MustGet("db").(*gorm.DB)
	//db.Raw("SELECT * FROM bicycles WHERE userid = ?", user.ID).Scan(&bicycles)
	db.Where("userid = ?", user.ID).Find(&bicycles)

	length := len(bicycles)
	serialized := make([]common.JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = bicycles[i].Serialize()
	}

	c.JSON(200, serialized)
}
