package constants

import (
	"os"
	"strconv"
	"strings"
)

var (
	HOST         string
	PORT         string
	DB_NAME      string
	DB_USER      string
	DB_PASS      string
	DB_HOST      string
	DB_PORT      string
	DB_IDLE_CONN int
	DB_MAX_CONN  int
	KAFKA_ACTIVE int
	BROKERS      []string
	TOPICS       []string
)

// setup env variables
func SetupEnv() {
	HOST = os.Getenv("HOST")
	PORT = os.Getenv("PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_IDLE_CONN, _ = strconv.Atoi(os.Getenv("DB_IDLE_CONN"))
	DB_MAX_CONN, _ = strconv.Atoi(os.Getenv("DB_MAX_CONN"))
	KAFKA_ACTIVE, _ = strconv.Atoi(os.Getenv("KAFKA_ACTIVE"))
	BROKERS = strings.Split(os.Getenv("BROKERS"), ",")
	TOPICS = strings.Split(os.Getenv("TOPICS"), ",")
}

const ()
