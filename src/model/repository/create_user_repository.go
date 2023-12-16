package repository

import (
	"context"
	"os"

	logger "github.com/vhtor/metaifrn-simulados-api/src/configuration/log"
	resterr "github.com/vhtor/metaifrn-simulados-api/src/configuration/rest_err"
	"github.com/vhtor/metaifrn-simulados-api/src/model"
	"github.com/vhtor/metaifrn-simulados-api/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (repository *userRepository) CreateUser(
	user model.UserInterface,
) (model.UserInterface, *resterr.RestErr) {
	
	logger.Info("Init createUser repository")
	
	collection_name := os.Getenv(MONGODB_USER_COLLECTION)
	collection := repository.databaseConnection.Collection(collection_name)

	userEntity := converter.DomainToEntity(user)

	result, err := collection.InsertOne(context.Background(), userEntity)
	if err != nil {
		return nil, resterr.NewInternalServerError(err.Error())
 	}

	userEntity.ID = result.InsertedID.(primitive.ObjectID)

	return converter.EntityToDomain(*userEntity), nil
}
