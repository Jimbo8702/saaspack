package types

import "time"

// booking can be make but is started as unfulfilled
//the booking then goes to the admin dashboard

type Booking struct {
	ID 			string    `bson:"id"         json:"id"`
	UserID 		string 	  `bson:"useriD"     json:"userID"`
	ProductID 	string 	  `bson:"productID"  json:"productID"`
	Length 		int 	  `bson:"length"     json:"length"`
	TotalPrice 	int 	  `bson:"totalPrice" json:"totalPrice"`
	Fulfilled 	bool 	  `bson:"fulfilled"  json:"fulfilled"`
	Extended 	bool 	  `bson:"extended"   json:"extended"`
	StartDate 	time.Time `bson:"startDate"  json:"startDate"`
	EndDate 	time.Time `bson:"endDate"    json:"endDate"`
	CreatedAt 	time.Time `bson:"createdAt"  json:"createdAt"`
}

type CreateBookingParams struct {
	UserID 		string 	  `json:"userID"`
	ProductID 	string 	  `json:"productID"`
	StartDate 	time.Time `json:"startDate"`
	EndDate 	time.Time `json:"endDate"`
}