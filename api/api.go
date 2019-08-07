package api

import (
	cont "tugas-akhir-2/controllers"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api/user")
	{
		api.POST("signup", cont.UserSignUp)
		api.POST("login", cont.UserLogin)
		api.GET("check", cont.Check)
		api.GET("retrieve", cont.UserRetrieve)
	}
}
