package main

import (
	"ecommerce-ums/cmd"
	"ecommerce-ums/helpers"
)

func main() {

	//load config
	helpers.SetupConfig()

	//load Log
	helpers.SetupLogger()

	//load DB
	helpers.SetupDB()

	//Load Redis
	// helpers.SetupRedis()

	// run kafka consumer
	// cmd.ServeKafkaConsumer()

	//run Http
	cmd.ServeHTTP()
}
