package main

import (
	"os"
	"typathon/config"
	"typathon/models/db"
	"typathon/routers"
)

var ConfigVariables config.ConfigType

func main() {
	ConfigVariables = *config.GetConfig()
	db.ConnectMongoDB(ConfigVariables)
	router := routers.InitRoute()
	port := os.Getenv("PORT")
	// port := utils.EnvVar("SERVER_PORT", "0.0.0.0:3000")
	router.Run(port)
}
