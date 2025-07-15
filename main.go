package main

import (
	"my-echo-framework/cmd"
	"my-echo-framework/helpers"
)

func main() {

	//load config
	helpers.SetupConfig()

	//load Log
	helpers.SetupLogger()

	//load DB
	helpers.SetupMySQL()

	//Load Redis
	// helpers.SetupRedis()

	// run kafka consumer
	// cmd.ServeKafkaConsumer()

	//run Http
	cmd.ServeHTTP()
}
