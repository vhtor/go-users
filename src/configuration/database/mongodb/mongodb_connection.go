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
	MONGODB_USER_DB_KEY = "MONGODB_USER_DB"
)

func NewMongoDBConnection(
	ctx context.Context,
) (*mongo.Database, error) {
	url := os.Getenv(MONGODB_URL_KEY)

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(url),
	)

	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Connected to MongoDB")
	return client.Database(MONGODB_USER_DB_KEY), nil
}
