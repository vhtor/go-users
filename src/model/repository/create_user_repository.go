package repository

import (
	"context"
	"os"

	logger "github.com/vhtor/metaifrn-simulados-api/src/configuration/log"
	resterr "github.com/vhtor/metaifrn-simulados-api/src/configuration/rest_err"
	"github.com/vhtor/metaifrn-simulados-api/src/model"
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

	userJSON, err := user.GetJSONValue(); 
	
	if err != nil {
		 return nil, resterr.NewInternalServerError(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), userJSON)
	if err != nil {
		return nil, resterr.NewInternalServerError(err.Error())
 	}

	user.SetID(result.InsertedID.(string))
	return user, nil
}
