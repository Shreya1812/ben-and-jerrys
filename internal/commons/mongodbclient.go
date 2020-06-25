package commons

import (
	"context"
	"fmt"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

func GetDatabase(config *configs.MongoDBConfig) (*mongo.Database, error) {

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", config.Host, config.Port))
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		return nil, ErrServerStartup{Err: errors.Wrap(err, "could not create mongoDB client")}
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	if err := client.Connect(ctx); err != nil {
		return nil, ErrServerStartup{Err: errors.Wrap(err, "could not connect to mongoDB client")}
	}

	ctx, _ = context.WithTimeout(context.Background(), 2*time.Second)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, ErrServerStartup{Err: errors.Wrap(err, "could not ping mongoDB client")}
	}

	db := client.Database(config.DatabaseName)

	if db == nil {
		return nil, ErrServerStartup{Err: fmt.Errorf("database in nil")}
	}

	return db, nil
}

func Disconnect(db *mongo.Database) error {

	if db != nil && db.Client() != nil {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		err := db.Client().Disconnect(ctx)

		if err != nil {
			return ErrServerShutdown{Err: errors.Wrap(err, "could not connect to mongoDB client")}
		}
	}

	return nil
}
