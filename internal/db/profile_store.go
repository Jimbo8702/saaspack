package db

import (
	"Jimbo8702/saaspack/internal/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const profileColl = "profile"

type ProfileStore interface {
	InsertProfile(context.Context, *types.Profile) (*types.Profile, error)
	GetProfileById(context.Context, string) (*types.Profile, error)
	ListProfiles(context.Context, Map) ([]*types.Profile, error)
	DeleteProfile(context.Context, string) error
	//update
}

type MongoProfileStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoProfileStore(client *mongo.Client) *MongoProfileStore {
	return &MongoProfileStore{
		client: client,
		coll: client.Database(DBNAME).Collection(profileColl),
	}
}

func (s *MongoProfileStore) InsertProfile(ctx context.Context, p *types.Profile) (*types.Profile, error) {
	res, err := s.coll.InsertOne(ctx, p)
	if err != nil {
		return nil, err
	}
	p.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return p, nil
}

func (s *MongoProfileStore) GetProfileById(ctx context.Context, id string) (*types.Profile, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var profile types.Profile
	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&profile); err != nil {
		return nil, err
	}
	return &profile, nil
}
 
func (s *MongoProfileStore) ListProfiles(ctx context.Context, filter Map) ([]*types.Profile, error) {
	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var profiles []*types.Profile
	if err := resp.All(ctx, &profiles); err != nil {
		return nil, err
	}
	return profiles, nil
}

func (s *MongoProfileStore) DeleteProfile(ctx context.Context, id string) error {
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