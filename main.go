package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	logger "github.com/vhtor/metaifrn-simulados-api/src/configuration/log"
	"github.com/vhtor/metaifrn-simulados-api/src/controller/routes"
)

func main() {

	logger.Info("Starting application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
