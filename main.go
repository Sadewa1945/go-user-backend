package main

import (
	"sadewadev/backend-api/config"
	"sadewadev/backend-api/database"
	"sadewadev/backend-api/routes"
)

func main (){

	config.LoadEnv()

	database.InitDB()
	
	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}