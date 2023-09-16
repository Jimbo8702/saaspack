package db

import (
	"Jimbo8702/saaspack/internal/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const categoryColl = "categories"

type CategoryStore interface {
	InsertCategory(context.Context, *types.Category) (*types.Category, error) 
	GetCategoryById(context.Context, string) (*types.Category, error)
	ListCategories(context.Context, Map) ([]*types.Category, error) 
	DeleteCategory(context.Context, string) error 
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

func (s MongoCategoryStore) GetCategoryById(ctx context.Context, id string) (*types.Category, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var category types.Category
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&category); err != nil {
		return nil, err
	}
	return &category, nil
}

func (s MongoCategoryStore) ListCategories(ctx context.Context, filter Map) ([]*types.Category, error) {
	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var categories []*types.Category
	if err := resp.All(ctx, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}

func (s MongoCategoryStore) DeleteCategory(ctx context.Context, id string) error {
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