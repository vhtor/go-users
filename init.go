package main

import (
	"context"
	"log"

	"github.com/vhtor/metaifrn-simulados-api/src/configuration/database/mongodb"
	user_controller "github.com/vhtor/metaifrn-simulados-api/src/controller/user"
	"github.com/vhtor/metaifrn-simulados-api/src/model/repository"
	"github.com/vhtor/metaifrn-simulados-api/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies() (user_controller.UserControllerInterface,) {

	database := initDatabase()

	repository := repository.NewUserRepository(database)
	service := service.NewUserService(repository)
	return user_controller.NewUserControllerInterface(service)
}

func initDatabase() *mongo.Database {
	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error trying to connect to database, error: %s", err.Error())
		return nil
	}

	return database
}