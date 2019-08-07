package main

import (
	"tugas-akhir-2/api"
	"tugas-akhir-2/database"
	"tugas-akhir-2/middlewares"

	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main() {
	// load .env environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	//initialize database
	db := database.CreateConnection()

	//app
	app := gin.Default()
	app.Use(database.Inject(db))
	app.Use(middlewares.JWTMiddleware())
	api.ApplyRoutes(app) //router
	app.Run(":3000")
}
