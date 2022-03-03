package main

import (
	"fmt"

	"github.com/OmarAouini/employee-manager/config"
	"github.com/OmarAouini/employee-manager/config/database"
	"github.com/OmarAouini/employee-manager/constants/roles"
	"github.com/sirupsen/logrus"
)

func main() {
	config.ConfigApp()
	database.ConnectDatabase()
	logrus.Info("ciao da docker container")
	fmt.Println(roles.IsPresent("admin"))
}
