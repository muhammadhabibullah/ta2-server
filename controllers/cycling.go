package controllers

import (
	"bytes"
	"strconv"
	"tugas-akhir-2/common"
	"tugas-akhir-2/middlewares"
	"tugas-akhir-2/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Cycling is alias for models.Cycling
type Cycling = models.Cycling

//CyclingCalendar get cycling data at requested month-year range
func CyclingCalendar(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	var cyclings []Cycling
	var b bytes.Buffer

	db := c.MustGet("db").(*gorm.DB)
	month := c.Param("m")
	year := c.Param("y")
	m, _ := strconv.Atoi(month)
	y, _ := strconv.Atoi(year)

	b.WriteString(year)
	b.WriteString("-")
	b.WriteString(month)
	b.WriteString("-")
	b.WriteString("01")
	datestart := b.String()
	b.Reset()

	b.WriteString(year)
	b.WriteString("-")
	b.WriteString(month)
	b.WriteString("-")
	switch m {
	case 1, 3, 5, 7, 8, 10, 12:
		b.WriteString("31")
	case 4, 6, 9, 11:
		b.WriteString("30")
	case 2:
		kabisat := (y%4 == 0)
		if kabisat {
			b.WriteString("29")
		} else {
			b.WriteString("28")
		}
	}
	datefinish := b.String()

	db.Where("userid = ? and starttime >= ? and starttime <= ?", user.ID, datestart, datefinish).Order("starttime desc").Find(&cyclings)

	length := len(cyclings)
	serialized := make([]common.JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = cyclings[i].Serialize()
	}

	c.JSON(200, serialized)

}

//CyclingRetrieve get requested-number lastest cycling of user (for list view UX)
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
