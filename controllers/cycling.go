package controllers

import (
	"strconv"
	"tugas-akhir-2/common"
	"tugas-akhir-2/middlewares"
	"tugas-akhir-2/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Cycling is alias for models.Cycling
type Cycling = models.Cycling

//CyclingRetrieve get 10 lastest cycling of user
func CyclingRetrieve(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	var cyclings []Cycling

	db := c.MustGet("db").(*gorm.DB)
	view := c.Param("view")
	page := c.Param("page")
	v, _ := strconv.Atoi(view)
	p, _ := strconv.Atoi(page)
	offset := (p - 1) * v
	//db.Raw("SELECT * FROM cyclings WHERE userid = ? ORDER BY starttime DESC;", user.ID).Scan(&cyclings)
	db.Limit(view).Offset(offset).Where("userid = ?", user.ID).Order("starttime desc").Find(&cyclings)

	length := len(cyclings)
	serialized := make([]common.JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = cyclings[i].Serialize()
	}

	c.JSON(200, serialized)
}
