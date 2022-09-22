package main

import (
	"typathon/config"
	"typathon/models/db"
	"typathon/routers"
	"typathon/utils"
)

var ConfigVariables config.ConfigType

func main() {
	ConfigVariables = *config.GetConfig()
	db.ConnectMongoDB(ConfigVariables)
	router := routers.InitRoute()
	// port := os.Getenv("PORT")
	port := utils.EnvVar("SERVER_PORT", ":80")
	router.Run(port)
}
