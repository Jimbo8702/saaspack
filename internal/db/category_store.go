package db

import (
	"Jimbo8702/saaspack/internal/types"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const categoryColl = "categories"

type CategoryStore interface {
	InsertCategory(context.Context, *types.Category) (*types.Category, error) 
}

type MongoCategoryStore struct {
	client 	*mongo.Client
	coll 	*mongo.Collection
}

func NewMongoCategoryStore(client *mongo.Client) *MongoCategoryStore {
	return &MongoCategoryStore{
		client: client,
		coll: client.Database(DBNAME).Collection(categoryColl),
	}
}

func (s MongoCategoryStore) InsertCategory(ctx context.Context, c *types.Category) (*types.Category, error) {
	res, err := s.coll.InsertOne(ctx, c)
	if err != nil {
		return nil, err
	}
	c.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return c, nil
}