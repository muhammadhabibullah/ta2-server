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

//CyclingRetrieve get 10 lastest cycling of user (for list view UX)
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

//CyclingGraph retrieve data for show cycling graph in certain time range
func CyclingGraph(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	var cyclings []Cycling

	db := c.MustGet("db").(*gorm.DB)
	y := c.Param("y")
	x := c.Param("x")
	//metric, _ := strconv.Atoi(view)
	//timerange, _ := strconv.Atoi(page)

	metric := ""
	switch y {
	case "D":
		metric = "distance"
	case "P":
		metric = "averagepace"
	case "E":
		metric = "elevationgain"
	case "HR":
		metric = "heartrate"
	case "C":
		metric = "calorieburned"
	case "T":
		metric = "finishtime"
	}
	switch x {
	case "1M":
		db.Select("id, starttime, ?", metric).Where("userid = ? and starttime >= DATE(NOW()) - INTERVAL 1 WEEK", user.ID).Order("starttime desc").Find(&cyclings)
	case "2M":
		db.Select("id, starttime, ?", metric).Where("userid = ? and starttime >= DATE(NOW()) - INTERVAL 2 WEEK", user.ID).Order("starttime desc").Find(&cyclings)
	case "1B":
		db.Select("id, starttime, ?", metric).Where("userid = ? and starttime >= DATE(NOW()) - INTERVAL 1 MONTH", user.ID).Order("starttime desc").Find(&cyclings)
	case "6B":
		db.Select("id, starttime, ?", metric).Where("userid = ? and starttime >= DATE(NOW()) - INTERVAL 6 MONTH", user.ID).Order("starttime desc").Find(&cyclings)
	case "1T":
		db.Select("id, starttime, ?", metric).Where("userid = ? and starttime >= DATE(NOW()) - INTERVAL 1 YEAR", user.ID).Order("starttime desc").Find(&cyclings)
	}

	length := len(cyclings)
	serialized := make([]common.JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = cyclings[i].Serialize()
	}

	c.JSON(200, serialized)
}

//CyclingDetail give detail of a cycling data
func CyclingDetail(c *gin.Context) {
	_ = middlewares.AuthorizedUser(c)

	var cycling Cycling

	db := c.MustGet("db").(*gorm.DB)
	cyclingid := c.Param("cyclingid")
	db.Where("id = ?", cyclingid).Find(&cycling)

	c.JSON(200, cycling.Serialize())
}
