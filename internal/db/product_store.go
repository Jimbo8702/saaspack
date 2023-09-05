package db

import (
	"Jimbo8702/saaspack/internal/types"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const productColl = "products"

type ProductStore interface {
	InsertProduct(context.Context, *types.Product) (*types.Product, error)
}

type MongoProductStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoProductStore(client *mongo.Client) *MongoProductStore {
	return &MongoProductStore{
		client: client,
		coll: client.Database(DBNAME).Collection(productColl),
	}
}

func (s *MongoProductStore) InsertProduct(ctx context.Context, product *types.Product) (*types.Product, error) {
	res, err := s.coll.InsertOne(ctx, product)
	if err != nil {
		return nil, err
	}
	product.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return product, nil
}