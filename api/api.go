package api

import (
	cont "tugas-akhir-2/controllers"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api/user")
	{
		api.POST("/signup", cont.UserSignUp)
		api.POST("/login", cont.UserLogin)
		api.GET("/renewtoken", cont.UserRenewToken)
		api.GET("/retrieve", cont.UserRetrieve)
	}
	api = r.Group("api/bicycle")
	{
		api.POST("/signup", cont.BicycleSignUp)
		api.GET("/retrieve", cont.BicycleRetrieve)
		api.PATCH("/edit/:id", cont.BicycleEdit)
		api.DELETE("/delete/:id", cont.BicycleDelete)
	}
	api = r.Group("api/target")
	{
		api.POST("/signup", cont.TargetSignUp)
		api.GET("/retrieve", cont.TargetRetrieve)
		api.GET("/lastestretrieve", cont.LastestTargetRetrieve)
		api.PATCH("/edit/:id", cont.TargetEdit)
		api.DELETE("/delete/:id", cont.TargetDelete)
	}
	api = r.Group("api/cycling")
	{
		api.GET("/calendar/:m/:y", cont.CyclingCalendar) 	   //Calendar View
		api.GET("/retrieve/:view/:page", cont.CyclingRetrieve) //List View
		api.GET("/graph/:y/:x/", cont.CyclingGraph)            //Graph
		api.GET("/detail/:cyclingid", cont.CyclingDetail)      //Detail in a Cycling
	}
}
