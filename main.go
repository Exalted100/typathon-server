package main

import (
	"fmt"
	"typathon/config"
	"typathon/models/db"
	"typathon/routers"
	"typathon/utils"
)

var ConfigVariables config.ConfigType

func main() {
	fmt.Println("Server is starting")
	ConfigVariables = *config.GetConfig()
	fmt.Println("Config variables have been set", ConfigVariables, "- Now connecting to the db")
	db.ConnectMongoDB(ConfigVariables)
	fmt.Println("Connected to the db successfully")
	router := routers.InitRoute()
	fmt.Println("Routes have been added. About to start running the server.")
	port := utils.EnvVar("SERVER_PORT", ":8080")
	router.Run(port)
}
