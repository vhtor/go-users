package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	logger "github.com/vhtor/metaifrn-simulados-api/src/configuration/log"
	"github.com/vhtor/metaifrn-simulados-api/src/controller/routes"
	user_controller "github.com/vhtor/metaifrn-simulados-api/src/controller/user"
	"github.com/vhtor/metaifrn-simulados-api/src/model/service"
)

func main() {

	logger.Info("Starting application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init dependencies
	service := service.NewUserService()
	userController := user_controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
