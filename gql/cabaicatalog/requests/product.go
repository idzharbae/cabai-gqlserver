package requests

import upload "github.com/smithaitufe/go-graphql-upload"

type ListProduct struct {
	Province  string
	Search    string
	Category  string
	OrderBy   string
	OrderType string
	Page      int32
	Limit     int32
}

type ProductsByShop struct {
	ShopID int32
}

type GetProduct struct {
	ID       int32
	SlugName string
}
type DeleteProduct struct {
	ID       int32
	SlugName string
	UserID   int32
}

type CreateProduct struct {
	ShopID      int32
	Name        string
	Quantity    int32
	PricePerKG  int32
	SlugName    string
	StockKG     float64
	Photo       *upload.GraphQLUpload
	Description string
	Category    string
}

type UpdateProduct struct {
	ID          int32
	ShopID      int32
	Name        string
	Quantity    int32
	PricePerKG  int32
	SlugName    string
	StockKG     float64
	Photo       *upload.GraphQLUpload
	Description string
	Category    string
}
