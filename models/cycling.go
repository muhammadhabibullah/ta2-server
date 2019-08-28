package models

import (
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
		"id":             c.ID,
		"starttime":      c.StartTime,
		"finishtime":     c.FinishTime,
		"averagepace":    c.AveragePace,
		"elevationgain":  c.ElevationGain,
		"distance":       c.Distance,
		"heartrate":      c.HeartRate,
		"calorieburned":  c.CalorieBurned,
		"percentofgoal":  c.PercentOfGoal,
		"recommendation": c.Recommendation,
		"userid":         c.UserID,
		"bicycleid":      c.BicycleID,
		"gpsid":          c.GPSID,
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
		"id":         c.ID,
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
			"id":         c.ID,
			"finishtime": c.FinishTime,
			"distance":   c.Distance,
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
			"id":        c.ID,
			"value":     c.Distance,
		}
	case "P":
		ret = common.JSON{
			"starttime": c.StartTime,
			"id":        c.ID,
			"value":     c.AveragePace,
		}
	case "E":
		ret = common.JSON{
			"starttime": c.StartTime,
			"id":        c.ID,
			"value":     c.ElevationGain,
		}
	case "HR":
		ret = common.JSON{
			"starttime": c.StartTime,
			"id":        c.ID,
			"value":     c.HeartRate,
		}
	case "C":
		ret = common.JSON{
			"starttime": c.StartTime,
			"id":        c.ID,
			"value":     c.CalorieBurned,
		}
	case "T":
		ret = common.JSON{
			"starttime": c.StartTime,
			"id":        c.ID,
			"value":     c.FinishTime,
		}
	}
	return ret
}
