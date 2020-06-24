package db

import (
	"context"
	"fmt"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/internal/data"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type iceCreamDBImpl struct {
	client *mongo.Client
	db     *mongo.Database
}

const collectionName = "icecream"

func New(config *configs.MongoDBConfig) IceCreamDB {
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%v:%v", config.Host, config.Port))
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(context.TODO()); err != nil {
		log.Fatalf("Cannot connect to mongodb: %v", err)
	}

	db := client.Database(config.DatabaseName)

	return &iceCreamDBImpl{
		client: client,
		db:     db,
	}
}

func (i *iceCreamDBImpl) Create(ctx context.Context, d *data.IceCream) (*data.IceCream, error) {
	c := i.db.Collection(collectionName)

	if _, err := c.InsertOne(ctx, d); err != nil {
		return nil, err // Insertion Error
	}
	return d, nil
}

func (i *iceCreamDBImpl) Update(ctx context.Context, d *data.IceCream) (*data.IceCream, error) {
	c := i.db.Collection(collectionName)

	filter := bson.D{{"product_id", d.ProductId}}
	opts := options.FindOneAndReplace().SetReturnDocument(options.After)
	result := &data.IceCream{}

	//res, err := c.UpdateOne(ctx, filter, d)
	err := c.FindOneAndReplace(ctx, filter, d, opts).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err // Item Not Found
		}
		return nil, err // Database error
	}

	return result, nil
}

func (i *iceCreamDBImpl) Delete(ctx context.Context, pId string) (*data.IceCream, error) {
	c := i.db.Collection(collectionName)

	filter := bson.D{{"product_id", pId}}
	result := &data.IceCream{}

	err := c.FindOneAndDelete(ctx, filter).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err // Item Not Found
		}
		return nil, err // Database error
	}
	return result, nil
}

func (i *iceCreamDBImpl) GetById(ctx context.Context, pId string) (*data.IceCream, error) {
	c := i.db.Collection(collectionName)

	filter := bson.D{{"product_id", pId}}
	result := &data.IceCream{}

	err := c.FindOne(ctx, filter).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err // Item Not Found
		}
		return nil, err // Database error
	}

	return result, nil
}

func (i *iceCreamDBImpl) Close() error {
	if i.client != nil {
		return i.client.Disconnect(context.TODO())
	}
	return nil
}
