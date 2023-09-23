package db

import (
	"Jimbo8702/saaspack/internal/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


const bookingColl = "bookings"

type BookingStore interface {
	InsertBooking(context.Context, *types.Booking) (*types.Booking, error)
	GetBookingById(context.Context, string) (*types.Booking, error)
	ListBookings(context.Context, Map) ([]*types.Booking, error)
	DeleteBooking(context.Context, string) error
	//update booking
}

type MongoBookingStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoBookingStore(client *mongo.Client) *MongoBookingStore {
	return &MongoBookingStore{
		client: client,
		coll: client.Database(DBNAME).Collection(bookingColl),
	}
}

func (s *MongoBookingStore) InsertBooking(ctx context.Context, b *types.Booking) (*types.Booking, error) {
	res, err := s.coll.InsertOne(ctx, b)
	if err != nil {
		return nil, err
	}
	b.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return b, nil
}

func (s *MongoBookingStore) GetBookingById(ctx context.Context, id string) (*types.Booking, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var booking types.Booking
	if err := s.coll.FindOne(ctx, bson.M{"_id:": oid}).Decode(&booking); err != nil {
		return nil, err
	}
	return &booking, nil
}

func (s *MongoBookingStore) ListBookings(ctx context.Context, filter Map) ([]*types.Booking, error) {
	resp, err := s.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var bookings []*types.Booking
	if err := resp.All(ctx, &bookings); err != nil {
		return nil, err
	}
	return bookings, nil
}

func (s *MongoBookingStore) DeleteBooking(ctx context.Context, id string) error {
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