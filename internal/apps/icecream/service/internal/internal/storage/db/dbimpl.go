package db

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/icecream/service/internal/internal/data"
	"github.com/Shreya1812/ben-and-jerrys/internal/commons"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/xerrors"
)

type iceCreamDBImpl struct {
	db *mongo.Database
}

const collectionName = "icecream"

func New(config *configs.MongoDBConfig) (IceCreamDB, error) {
	db, err := commons.GetDatabase(config)

	if err != nil {
		return nil, err
	}

	return &iceCreamDBImpl{
		db: db,
	}, nil
}

func (i *iceCreamDBImpl) Create(ctx context.Context, d *data.IceCream) (*data.IceCream, error) {
	c := i.db.Collection(collectionName)

	_, err := i.GetByProductId(ctx, d.ProductId)

	if err != nil {
		if xerrors.Is(err, commons.ErrItemNotFound) {
			if _, err := c.InsertOne(ctx, d); err != nil {
				return nil, err // Insertion Error
			}
			return d, nil
		}
		return nil, err
	}

	return nil, commons.ErrItemAlreadyExists
}

func (i *iceCreamDBImpl) Update(ctx context.Context, d *data.IceCream) (*data.IceCream, error) {
	c := i.db.Collection(collectionName)

	filter := bson.D{{"product_id", d.ProductId}}
	opts := options.FindOneAndReplace().SetReturnDocument(options.After)
	result := &data.IceCream{}

	err := c.FindOneAndReplace(ctx, filter, d, opts).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, commons.ErrItemNotFound
		}
		return nil, err // Database error
	}

	return result, nil
}

func (i *iceCreamDBImpl) DeleteByProductId(ctx context.Context, pId string) (*data.IceCream, error) {
	c := i.db.Collection(collectionName)

	filter := bson.D{{"product_id", pId}}
	result := &data.IceCream{}

	err := c.FindOneAndDelete(ctx, filter).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, commons.ErrItemNotFound
		}
		return nil, err // Database error
	}
	return result, nil
}

func (i *iceCreamDBImpl) GetByProductId(ctx context.Context, pId string) (*data.IceCream, error) {
	c := i.db.Collection(collectionName)

	filter := bson.D{{"product_id", pId}}
	result := &data.IceCream{}

	err := c.FindOne(ctx, filter).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, commons.ErrItemNotFound
		}
		return nil, err // Database error
	}

	return result, nil
}

func (i *iceCreamDBImpl) GetList(ctx context.Context, searchOptions *data.IceCreamSearchOptions) (*data.IceCreamListResult, error) {
	c := i.db.Collection(collectionName)

	filter := bson.M{}
	if searchOptions.LastId != "" {
		objectID, _ := primitive.ObjectIDFromHex(searchOptions.LastId)
		filter = bson.M{"_id": bson.M{"$lt": objectID}}
	}

	opts := options.Find().SetSort(bson.D{{"_id", -1}}).SetLimit(searchOptions.Limit)

	cursor, err := c.Find(ctx, filter, opts)

	if err != nil {
		return nil, err
	}

	var results []*data.IceCream
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	lastId := ""

	if len(results) > 0 {
		lastId = results[len(results)-1].Id.Hex()
	}

	return &data.IceCreamListResult{
		IceCreams: results,
		LastId:    lastId,
	}, nil
}

func (i *iceCreamDBImpl) Close() error {
	return commons.Disconnect(i.db)
}
