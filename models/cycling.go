package models

import (
	"tugas-akhir-2/common"
)

//Cycling model
type Cycling struct {
	ID            uint    `gorm:"column:id;primary_key"`
	StartTime     string  `gorm:"column:starttime;type:timestamp; not null"`
	FinishTime    string  `gorm:"column:finishtime;type:timestamp; not null"`
	AveragePace   float64 `gorm:"column:averagepace;type:float; not null"`
	ElevationGain float64 `gorm:"column:elevationgain;type:float; not null"`
	Distance      float64 `gorm:"column:distance;type:float; not null"`
	HeartRate     uint    `gorm:"column:heartrate; not null"`
	CalorieBurned float64 `gorm:"column:calorieburned;type:float; not null"`
	PercentOfGoal float64 `gorm:"column:percentofgoal;type:float; not null"`
	UserID        uint    `gorm:"column:userid";not null`
	BicycleID     uint    `gorm:"column:bicycleid";not null`
	GPSID         uint    `gorm:"column:gpsid";not null`
}

// Serialize serializes bicycle data
func (c *Cycling) Serialize() common.JSON {
	return common.JSON{
		"id":            c.ID,
		"starttime":     c.StartTime,
		"finishtime":    c.FinishTime,
		"averagepace":   c.AveragePace,
		"elevationgain": c.ElevationGain,
		"distance":      c.Distance,
		"heartrate":     c.HeartRate,
		"calorieburned": c.CalorieBurned,
		"percentofgoal": c.PercentOfGoal,
		"userid":        c.UserID,
		"bicycleid":     c.BicycleID,
		"gpsid":         c.GPSID,
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
	c.UserID = uint(m["userid"].(float64))
	c.BicycleID = uint(m["bicycleid"].(float64))
	c.GPSID = uint(m["gpsid"].(float64))
}
