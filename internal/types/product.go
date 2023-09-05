package types

type Category struct {
	ID 	  string `bson:"id" json:"id"`
	Value string `bson:"value" json:"value"`
}

func NewCategory(v string) *Category {
	return &Category{
		Value: v,
	}
}

type Product struct {
	ID 			string 	 `bson:"id"          json:"id"`
	CategoryID 	string 	 `bson:"categoryID"  json:"category"`
	Name 		string 	 `bson:"name"        json:"name"`
	Description string	 `bson:"description" json:"description"`
	Price 		string	 `bson:"price"       json:"price"`
	Deposit 	string	 `bson:"deposit"     json:"deposit"`
	RentPeriod 	string 	 `bson:"period"      json:"period"`
}

type CreateProductParams struct {
	CategoryID 	string 	 `json:"categoryID"  params:"required"`
	Name 		string 	 `json:"name"        params:"required"`
	Description string	 `json:"description" params:"required"`
	Price 		string	 `json:"price"       params:"required"`
	Deposit 	string	 `json:"deposit"     params:"required"`
	RentPeriod 	string 	 `json:"period"      params:"required"`
}

func NewProductFromParams(params CreateProductParams) *Product {
	return &Product{
		CategoryID: params.CategoryID,
		Name: params.Name,
		Description: params.Description,
		Price: params.Price,
		Deposit: params.Deposit,
		RentPeriod: params.RentPeriod,
	}
}