package requests

import upload "github.com/smithaitufe/go-graphql-upload"

type ListProduct struct {
	ShopID int32
	Page   int32
	Limit  int32
}

type GetProduct struct {
	ID       int32
	SlugName string
}

type CreateProduct struct {
	ShopID     int32
	Name       string
	Quantity   int32
	PricePerKG int32
	SlugName   string
	StockKG    float64
	Photo      upload.GraphQLUpload
}

type UpdateProduct struct {
	ID         int32
	ShopID     int32
	Name       string
	Quantity   int32
	PricePerKG int32
	SlugName   string
	StockKG    float64
	Photo      upload.GraphQLUpload
}
