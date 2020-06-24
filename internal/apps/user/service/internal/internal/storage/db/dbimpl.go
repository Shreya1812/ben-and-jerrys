package db

import (
	"context"
	"fmt"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/internal/internal/data"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type userDBImpl struct {
	client *mongo.Client
	db     *mongo.Database
}

const collectionName = "user"

func New(config *configs.MongoDBConfig) UserDB {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", config.Host, config.Port))
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatalf("Cannot connect to mongodb: %v", err)
	}

	db := client.Database(config.DatabaseName)

	return &userDBImpl{
		client: client,
		db:     db,
	}
}

func (u userDBImpl) Create(ctx context.Context, d *data.User) error {
	c := u.db.Collection(collectionName)

	if _, err := c.InsertOne(ctx, d); err != nil {
		return err // Insertion Error
	}
	return nil
}

func (u userDBImpl) Update(ctx context.Context, d *data.User) error {
	c := u.db.Collection(collectionName)

	filter := bson.D{{"email", d.Email}}
	res, err := c.ReplaceOne(ctx, filter, d)

	if err != nil {
		return err // Database error
	}
	if res.MatchedCount == 0 {
		return err // Item Not Found
	}

	return nil
}

func (u userDBImpl) Delete(ctx context.Context, email string) error {
	c := u.db.Collection(collectionName)

	filter := bson.D{{"email", email}}
	res, err := c.DeleteOne(ctx, filter)

	if err != nil {
		return err // Database error
	}
	if res.DeletedCount == 0 {
		return err // Item Not Found
	}

	return nil
}

func (u userDBImpl) GetByEmail(ctx context.Context, email string) (*data.User, error) {
	c := u.db.Collection(collectionName)

	filter := bson.D{{"email", email}}
	result := &data.User{}

	err := c.FindOne(ctx, filter).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err // Item Not Found
		}

		return nil, err // Database Error
	}

	return result, nil
}

func (u userDBImpl) Close() error {
	if u.client != nil {
		return u.client.Disconnect(context.TODO())
	}
	return nil
}
