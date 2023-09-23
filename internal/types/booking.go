package types

import (
	"Jimbo8702/saaspack/util"
	"time"
)

// booking can be make but is started as unfulfilled
//the booking then goes to the admin dashboard

type Booking struct {
	ID 			string    `bson:"id"         json:"id"`
	UserID 		string 	  `bson:"userID"     json:"userID"`
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
	UserID 		string 	  `json:"userID"     params:"required"`
	ProductID 	string 	  `json:"productID"  params:"required"`
	TotalPrice 	int 	  `json:"totalPrice" params:"required"`
	StartDate 	time.Time `json:"startDate"  params:"required"`
	EndDate 	time.Time `json:"endDate"    params:"required"`
}

func NewBookingFromParams(params CreateBookingParams) *Booking {
	length := util.DaysBetween(params.StartDate, params.EndDate)

	return &Booking{
		UserID: 	params.UserID,
		ProductID: 	params.ProductID,
		Length: 	length,
		TotalPrice: params.TotalPrice,
		Fulfilled: 	false,
		Extended: 	false,
		StartDate: 	params.StartDate,
		EndDate: 	params.EndDate,
		CreatedAt: 	time.Now(),
	}
}