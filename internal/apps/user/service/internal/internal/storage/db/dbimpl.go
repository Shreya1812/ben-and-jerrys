package db

import (
	"context"
	"github.com/Shreya1812/ben-and-jerrys/internal/apps/user/service/internal/internal/data"
	"github.com/Shreya1812/ben-and-jerrys/internal/commons"
	"github.com/Shreya1812/ben-and-jerrys/internal/configs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/xerrors"
)

type userDBImpl struct {
	db *mongo.Database
}

const collectionName = "user"

func New(config *configs.MongoDBConfig) (UserDB, error) {
	db, err := commons.GetDatabase(config)

	if err != nil {
		return nil, err
	}

	return &userDBImpl{
		db: db,
	}, nil
}

func (u *userDBImpl) Create(ctx context.Context, d *data.User) error {
	c := u.db.Collection(collectionName)

	_, err := u.GetByEmail(ctx, d.Email)

	if err != nil {
		if xerrors.Is(err, commons.ErrItemNotFound) {
			if _, err := c.InsertOne(ctx, d); err != nil {
				return err // Insertion Error
			}
			return nil
		}
		return err // Database error
	}

	return commons.ErrItemAlreadyExists
}

func (u *userDBImpl) Update(ctx context.Context, d *data.User) error {
	c := u.db.Collection(collectionName)

	filter := bson.D{{"email", d.Email}}
	res, err := c.ReplaceOne(ctx, filter, d)

	if err != nil {
		return err // Database error
	}
	if res.MatchedCount == 0 {
		return commons.ErrItemNotFound
	}

	return nil
}

func (u *userDBImpl) Delete(ctx context.Context, email string) error {
	c := u.db.Collection(collectionName)

	filter := bson.D{{"email", email}}
	res, err := c.DeleteOne(ctx, filter)

	if err != nil {
		return err // Database error
	}
	if res.DeletedCount == 0 {
		return commons.ErrItemNotFound
	}

	return nil
}

func (u *userDBImpl) GetByEmail(ctx context.Context, email string) (*data.User, error) {
	c := u.db.Collection(collectionName)

	filter := bson.D{{"email", email}}
	result := &data.User{}

	err := c.FindOne(ctx, filter).Decode(result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, commons.ErrItemNotFound
		}

		return nil, err
	}

	return result, nil
}

func (u *userDBImpl) Close() error {
	return commons.Disconnect(u.db)
}
