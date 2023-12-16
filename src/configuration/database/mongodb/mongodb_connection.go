package mongodb

import (
	"context"
	"os"

	logger "github.com/vhtor/metaifrn-simulados-api/src/configuration/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MONGODB_URL_KEY = "MONGODB_URL"
	MONGODB_USERS_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	
	mongodb_uri := os.Getenv(MONGODB_URL_KEY)
	mongodb_database := os.Getenv(MONGODB_USERS_DB)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongodb_uri),
	)

	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Connected to MongoDB")
	return client.Database(mongodb_database), nil
}
