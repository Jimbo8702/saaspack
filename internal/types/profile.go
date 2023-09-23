package types

import "time"


type BillingInfo struct {
	AddressOne 	string `json:"addressOne"`
	AddressTwo 	string `json:"addressTwo"`
	City		string `json:"city"`
	State		string `json:"state"`
	Zip 		int	   `json:"zip"`
}

type Profile struct {
	BillingInfo
	ID 		  string 	`json:"id"`
	AuthID 	  string    `json:"-"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email 	  string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateProfileParams struct {
	AuthID 	  string    `json:"-"          params:"required"`
	FirstName string    `json:"firstName"  params:"required"`
	LastName  string    `json:"lastName"   params:"required"`
	Email 	  string    `json:"email"      params:"required"`
	CreatedAt time.Time `json:"createdAt"  params:"required"`
}

func NewProfileFromParams(params CreateProfileParams) *Profile {
	return &Profile{
		AuthID: 	 params.AuthID,
		BillingInfo: BillingInfo{},
		FirstName: 	 params.FirstName,
		LastName: 	 params.LastName,
		Email: 		 params.Email,
	}
}