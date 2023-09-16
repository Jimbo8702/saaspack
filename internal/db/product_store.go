package db

import (
	"Jimbo8702/saaspack/internal/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const productColl = "products"

type ProductStore interface {
	InsertProduct(context.Context, *types.Product) (*types.Product, error)
	GetProductById(context.Context, string) (*types.Product, error)
	ListProducts(context.Context, Map) ([]*types.Product, error) 
	DeleteProduct(context.Context, string) error 
	//add update
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

func (s *MongoProductStore) GetProductById(ctx context.Context, id string) (*types.Product, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var product types.Product
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&product); err != nil {
		return nil, err
	}
	return &product, nil
}

func (s *MongoProductStore) ListProducts(ctx context.Context, filter Map) ([]*types.Product, error) {
	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var products []*types.Product
	if err := resp.All(ctx, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func (s *MongoProductStore) DeleteProduct(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}