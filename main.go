package main

import (
	"fmt"

	"github.com/OmarAouini/employee-manager/config"
	"github.com/OmarAouini/employee-manager/config/database"
	"github.com/OmarAouini/employee-manager/web"
)

func main() {
	config.ConfigApp()
	database.ConnectDatabase()
	printBanner()
	web.RunServer()
}

func printBanner() {
	fmt.Println("2022 OmarAouini")
	fmt.Println("EMPLOYEE-MANAGER")
	fmt.Printf("IS_PRODUCTION: %v\n", config.CONFIG.IsProduction)
	fmt.Println("======================")
}
