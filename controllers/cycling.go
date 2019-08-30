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

//CyclingCalendarMonth get cycling data at requested month-year range
func CyclingCalendarMonth(c *gin.Context) {
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
		serialized[i] = cyclings[i].CustomCalendarSerialize("M")
	}

	c.JSON(200, serialized)

}

//CyclingCalendarDate get cycling data at requested date
func CyclingCalendarDate(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	var cyclings []Cycling
	var b bytes.Buffer

	db := c.MustGet("db").(*gorm.DB)
	day := c.Param("d")
	month := c.Param("m")
	year := c.Param("y")

	b.WriteString(year)
	b.WriteString("-")
	b.WriteString(month)
	b.WriteString("-")
	b.WriteString(day)
	b.WriteString("T00:00:00+07:00")
	datestart := b.String()
	b.Reset()

	b.WriteString(year)
	b.WriteString("-")
	b.WriteString(month)
	b.WriteString("-")
	b.WriteString(day)
	b.WriteString("T23:59:59+07:00")
	datefinish := b.String()

	db.Where("userid = ? and starttime >= ? and starttime <= ?", user.ID, datestart, datefinish).Order("starttime desc").Find(&cyclings)

	length := len(cyclings)
	serialized := make([]common.JSON, length, length)

	for i := 0; i < length; i++ {
		serialized[i] = cyclings[i].CustomCalendarSerialize("D")
	}

	c.JSON(200, serialized)

}

//CyclingList get requested-number lastest cycling of user (for list view UX)
func CyclingList(c *gin.Context) {
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
		serialized[i] = cyclings[i].CustomListSerialize()
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
		db.Select([]string{"id", "starttime", metric}).Where("userid = ? and starttime >= DATE(NOW()) - INTERVAL 1 WEEK", user.ID).Order("starttime desc").Find(&cyclings)
	case "2M":
		db.Select([]string{"id", "starttime", metric}).Where("userid = ? and starttime >= DATE(NOW()) - INTERVAL 2 WEEK", user.ID).Order("starttime desc").Find(&cyclings)
	case "1B":
		db.Select([]string{"id", "starttime", metric}).Where("userid = ? and starttime >= DATE(NOW()) - INTERVAL 1 MONTH", user.ID).Order("starttime desc").Find(&cyclings)
	case "6B":
		db.Select([]string{"id", "starttime", metric}).Where("userid = ? and starttime >= DATE(NOW()) - INTERVAL 6 MONTH", user.ID).Order("starttime desc").Find(&cyclings)
	case "1T":
		db.Select([]string{"id", "starttime", metric}).Where("userid = ? and starttime >= DATE(NOW()) - INTERVAL 1 YEAR", user.ID).Order("starttime desc").Find(&cyclings)
	}

	length := len(cyclings)
	// serialized := make([]common.JSON, length, length)

	// for i := 0; i < length; i++ {
	// 	serialized[i] = cyclings[i].Serialize()
	// }

	serializedSelected := make([]common.JSON, length, length)
	switch y {
	case "D":
		// metric = "distance"
		for i := 0; i < length; i++ {
			serializedSelected[i] = cyclings[i].CustomGraphSerialize("D")
		}
	case "P":
		// metric = "averagepace"
		for i := 0; i < length; i++ {
			serializedSelected[i] = cyclings[i].CustomGraphSerialize("P")
		}
	case "E":
		// metric = "elevationgain"
		for i := 0; i < length; i++ {
			serializedSelected[i] = cyclings[i].CustomGraphSerialize("E")
		}
	case "HR":
		// metric = "heartrate"
		for i := 0; i < length; i++ {
			serializedSelected[i] = cyclings[i].CustomGraphSerialize("HR")
		}
	case "C":
		// metric = "calorieburned"
		for i := 0; i < length; i++ {
			serializedSelected[i] = cyclings[i].CustomGraphSerialize("C")
		}
	case "T":
		// metric = "finishtime"
		for i := 0; i < length; i++ {
			serializedSelected[i] = cyclings[i].CustomGraphSerialize("T")
		}
	}

	c.JSON(200, serializedSelected)
}

//CyclingDetail give detail of a cycling data
func CyclingDetail(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	var cycling Cycling

	db := c.MustGet("db").(*gorm.DB)
	cyclingid := c.Param("cyclingid")
	db.Where("userid = ? and id = ?", user.ID, cyclingid).Find(&cycling)

	if cycling.UserID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	c.JSON(200, cycling.Serialize())
}

//CyclingProgress return code for draw arrow in cycling data
func CyclingProgress(c *gin.Context) {
	user := middlewares.AuthorizedUser(c)

	var b bytes.Buffer

	var cycling Cycling
	var lastcycling Cycling
	db := c.MustGet("db").(*gorm.DB)
	cyclingid := c.Param("cyclingid")

	db.Where("userid = ? and id = ?", user.ID, cyclingid).Find(&cycling)

	if cycling.UserID != user.ID {
		c.AbortWithStatus(403)
		return
	}

	startTime := cycling.StartTime[:len(cycling.StartTime)-6] //cut +07:00 char
	b.WriteString(startTime)
	b.WriteString(".000Z") // add .000Z char
	sT := b.String()
	//fmt.Println(sT)

	db.Limit(1).Where("userid = ? and starttime < ?", user.ID, sT).Order("starttime desc").Find(&lastcycling)

	//Compare
	p := "0"   //pace
	e := "0"   //elevation
	d := "0"   //distance
	k := "0"   //kalori
	h := "0"   //heartrate
	per := "0" //percent

	//check if lastcycling exists //not required bcs lastcycling is null and if its compared it'll be true
	if cycling.AveragePace > lastcycling.AveragePace {
		p = "1"
	} else if cycling.AveragePace < lastcycling.AveragePace {
		p = "-1"
	}

	if cycling.ElevationGain > lastcycling.ElevationGain {
		e = "1"
	} else if cycling.ElevationGain < lastcycling.ElevationGain {
		e = "-1"
	}

	if cycling.Distance > lastcycling.Distance {
		d = "1"
	} else if cycling.Distance < lastcycling.Distance {
		d = "-1"
	}

	if cycling.CalorieBurned > lastcycling.CalorieBurned {
		k = "1"
	} else if cycling.CalorieBurned < lastcycling.CalorieBurned {
		k = "-1"
	}

	//Heart Rate
	if cycling.HeartRate > lastcycling.HeartRate {
		if float32(cycling.HeartRate) > (float32((220 - user.CountAge())) * 0.85) {
			h = "2"
		} else {
			h = "1"
		}
	} else if cycling.HeartRate < lastcycling.HeartRate {
		if float32(cycling.HeartRate) > (float32((220 - user.CountAge())) * 0.85) {
			h = "-2"
		} else {
			h = "-1"
		}
	}

	//Percent
	if cycling.PercentOfGoal > lastcycling.PercentOfGoal {
		per = "1"
	} else if cycling.PercentOfGoal < lastcycling.PercentOfGoal {
		per = "-1"
	}

	c.JSON(200, common.JSON{
		"averagepace":   p,
		"elevationgain": e,
		"distance":      d,
		"calorieburned": k,
		"heartrate":     h,
		"percentofgoal": per,
	})

}
