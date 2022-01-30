package main

import (
	"cloudCharts/config"
	"cloudCharts/routes"
	"fmt"
)

func main() {
	config.ConnectDatabase()
	r := routes.SetupRoutes()
	err := r.Run("0.0.0.0:8002")
	if err != nil {
		fmt.Println(err.Error())
	}
}
