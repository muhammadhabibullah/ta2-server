package models

import (
	"fmt"
	"tugas-akhir-2/common"
)

//Cycling model
type Cycling struct {
	ID             uint    `gorm:"column:id;primary_key"`
	StartTime      string  `gorm:"column:starttime;type:timestamp; not null"`
	FinishTime     string  `gorm:"column:finishtime;type:timestamp; not null"`
	AveragePace    float64 `gorm:"column:averagepace;type:float; not null"`
	ElevationGain  float64 `gorm:"column:elevationgain;type:float; not null"`
	Distance       float64 `gorm:"column:distance;type:float; not null"`
	HeartRate      uint    `gorm:"column:heartrate; not null"`
	CalorieBurned  float64 `gorm:"column:calorieburned;type:float; not null"`
	PercentOfGoal  float64 `gorm:"column:percentofgoal;type:float; not null"`
	Recommendation string  `gorm:"column:recommendation;type:text; not null"`
	UserID         uint    `gorm:"column:userid;not null"`
	BicycleID      uint    `gorm:"column:bicycleid;not null"`
	GPSID          uint    `gorm:"column:gpsid;not null"`
}

// Serialize serializes bicycle data
func (c *Cycling) Serialize() common.JSON {
	return common.JSON{
		"id":             fmt.Sprint(c.ID),
		"starttime":      c.StartTime,
		"finishtime":     c.FinishTime,
		"averagepace":    fmt.Sprint(c.AveragePace),
		"elevationgain":  fmt.Sprint(c.ElevationGain),
		"distance":       fmt.Sprint(c.Distance),
		"heartrate":      fmt.Sprint(c.HeartRate),
		"calorieburned":  fmt.Sprint(c.CalorieBurned),
		"percentofgoal":  fmt.Sprint(c.PercentOfGoal),
		"recommendation": c.Recommendation,
		"userid":         fmt.Sprint(c.UserID),
		"bicycleid":      fmt.Sprint(c.BicycleID),
		"gpsid":          fmt.Sprint(c.GPSID),
	}
}

func (c *Cycling) Read(m common.JSON) {
	c.ID = uint(m["id"].(float64))
	c.StartTime = m["starttime"].(string)
	c.FinishTime = m["finishtime"].(string)
	c.AveragePace = m["averagepace"].(float64)
	c.ElevationGain = m["elevationgain"].(float64)
	c.Distance = m["distance"].(float64)
	c.HeartRate = uint(m["heartrate"].(float64))
	c.CalorieBurned = m["calorieburned"].(float64)
	c.PercentOfGoal = m["percentofgoal"].(float64)
	c.Recommendation = m["recommendation"].(string)
	c.UserID = uint(m["userid"].(float64))
	c.BicycleID = uint(m["bicycleid"].(float64))
	c.GPSID = uint(m["gpsid"].(float64))
}

// CustomListSerialize serializes cycling data
func (c *Cycling) CustomListSerialize() common.JSON {
	return common.JSON{
		"starttime":  c.StartTime,
		"id":         fmt.Sprint(c.ID),
		"finishtime": c.FinishTime,
	}

}

// CustomCalendarSerialize serializes cycling data
func (c *Cycling) CustomCalendarSerialize(on string) common.JSON {
	ret := common.JSON{}
	switch on {
	case "M": //month
		ret = common.JSON{
			"starttime": c.StartTime,
		}
	case "D": //date
		ret = common.JSON{
			"starttime":  c.StartTime,
			"id":         fmt.Sprint(c.ID),
			"finishtime": c.FinishTime,
			"distance":   fmt.Sprint(c.Distance),
		}
	}
	return ret
}

// CustomGraphSerialize serializes cycling data based on parameter
func (c *Cycling) CustomGraphSerialize(metric string) common.JSON {
	ret := common.JSON{}
	switch metric {
	case "D":
		ret = common.JSON{
			"starttime": c.StartTime,
			"id":        fmt.Sprint(c.ID),
			"value":     fmt.Sprint(c.Distance),
		}
	case "P":
		ret = common.JSON{
			"starttime": c.StartTime,
			"id":        fmt.Sprint(c.ID),
			"value":     fmt.Sprint(c.AveragePace),
		}
	case "E":
		ret = common.JSON{
			"starttime": c.StartTime,
			"id":        fmt.Sprint(c.ID),
			"value":     fmt.Sprint(c.ElevationGain),
		}
	case "HR":
		ret = common.JSON{
			"starttime": c.StartTime,
			"id":        fmt.Sprint(c.ID),
			"value":     fmt.Sprint(c.HeartRate),
		}
	case "C":
		ret = common.JSON{
			"starttime": c.StartTime,
			"id":        fmt.Sprint(c.ID),
			"value":     fmt.Sprint(c.CalorieBurned),
		}
	case "T":
		ret = common.JSON{
			"starttime": c.StartTime,
			"id":        fmt.Sprint(c.ID),
			"value":     fmt.Sprint(c.FinishTime),
		}
	}
	return ret
}

//GPSRawData type
type GPSRawData struct {
	ID        uint    `gorm:"column:id;primary_key"`
	CyclingID uint    `gorm:"column:cyclindid;not null"`
	Second    uint    `gorm:"column:s;not null"`
	Lat       float64 `gorm:"column:lt;not null"`
	Long      float64 `gorm:"column:lg;not null"`
}

//Serialize func
func (g *GPSRawData) Serialize() common.JSON {
	return common.JSON{
		"seconds": fmt.Sprint(g.Second),
		"lat":     fmt.Sprint(g.Lat),
		"long":    fmt.Sprint(g.Long),
	}
}

//CyclingRawData type
type CyclingRawData struct {
	ID        uint    `gorm:"column:id;primary_key"`
	Received  string  `gorm:"column:r;type:timestamp; not null"`
	BicycleID uint    `gorm:"column:b; not null"`
	When      string  `gorm:"column:w; not null"`
	Lat       float64 `gorm:"column:t; not null"`
	Long      float64 `gorm:"column:g; not null"`
	Alt       float64 `gorm:"column:a; not null"`
	Pace      float64 `gorm:"column:p; not null"`
	Distance  float64 `gorm:"column:d; not null"`
	Elevation float64 `gorm:"column:e; not null"`
	HeartRate uint    `gorm:"column:h; not null"`
	Seconds   uint    `gorm:"column:s; not null"`
}

// Serialize serializes bicycle data
func (c *CyclingRawData) Serialize() common.JSON {
	return common.JSON{
		"id":        c.ID,
		"received":  c.Received,
		"bicycleid": c.BicycleID,
		"when":      c.When,
		"lat":       c.Lat,
		"long":      c.Long,
		"alt":       c.Alt,
		"pace":      c.Pace,
		"distance":  c.Distance,
		"elevation": c.Elevation,
		"heartrate": c.HeartRate,
		"seconds":   c.Seconds,
	}
}
