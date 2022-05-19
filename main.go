package main

import (
	"fmt"
	"log"
	"os"

	"github.com/OmarAouini/employee-manager/api"
	"github.com/OmarAouini/employee-manager/constants"
	"github.com/OmarAouini/employee-manager/core"
	"github.com/OmarAouini/employee-manager/kafka"
	"github.com/OmarAouini/employee-manager/utils"
)

func main() {
	core.LoadDotEnv()
	constants.SetupEnv()
	core.ConnectDb(constants.DB_USER, constants.DB_PASS, constants.DB_HOST, constants.DB_PORT, constants.DB_NAME, constants.DB_IDLE_CONN, constants.DB_MAX_CONN)
	core.Migrate()
	core := core.CoreV1{}
	server := api.Server{Core: &core}
	server.Init()
	utils.PrintAppInfo("employee-manager-api", os.Getenv("APP_ENV"), constants.DB_NAME, constants.DB_USER, constants.DB_HOST, constants.DB_PORT, constants.DB_IDLE_CONN, constants.DB_MAX_CONN, constants.HOST, constants.PORT, constants.BROKERS, constants.TOPICS)

	// if kafka feature active, init kafkaConsumer and subscribe
	if constants.KAFKA_ACTIVE == 1 {
		//parallel goroutine for subscribe to topics in order to trigger actions
		KafkaConsumerGroup := kafka.MsgConsumerGroup{
			Core:   &core,
			Topics: constants.TOPICS,
		}
		go func() {
			kafka.SubscribeTopics(KafkaConsumerGroup)
		}()
	}

	//panic logger to file
	defer func() {
		if x := recover(); x != nil {
			// recovering from a panic; x contains whatever was passed to panic()
			msg := fmt.Sprintf("run time panic: %v\n", x)
			log.Print(msg)

			f, err := os.OpenFile("panic_log.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
			if err != nil {
				panic(err)
			}

			defer f.Close()

			if _, err = f.WriteString(msg); err != nil {
				panic(err)
			}

			panic(x)
		}
	}()

	//run
	server.Run()
}
