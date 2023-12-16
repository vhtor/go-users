package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vhtor/metaifrn-simulados-api/src/configuration/database/mongodb"
	logger "github.com/vhtor/metaifrn-simulados-api/src/configuration/log"
	"github.com/vhtor/metaifrn-simulados-api/src/controller/routes"
	user_controller "github.com/vhtor/metaifrn-simulados-api/src/controller/user"
	"github.com/vhtor/metaifrn-simulados-api/src/model/repository"
	"github.com/vhtor/metaifrn-simulados-api/src/model/service"
)

func main() {

	logger.Info("Starting application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init database connection
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error: %s", err.Error())
		return
	}

	// Init dependencies
	repository := repository.NewUserRepository(database)
	service := service.NewUserService(repository)
	userController := user_controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
