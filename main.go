package main

import (
	"typathon/config"
	"typathon/routers"
	"typathon/utils"
	"typathon/models/db"
)

var ConfigVariables config.ConfigType

func main() {
	ConfigVariables = *config.GetConfig()
	db.ConnectMongoDB(ConfigVariables)
	router := routers.InitRoute()
	port := utils.EnvVar("SERVER_PORT", ":80")
	router.Run(port)
}
