package requests

type ListProduct struct {
	Page  int32
	Limit int32
}

type CreateProduct struct {
	ShopID     int32
	Name       string
	Quantity   int32
	PricePerKG int32
	SlugName   string
	StockKG    float64
}

type UpdateProduct struct {
	ID         int32
	ShopID     int32
	Name       string
	Quantity   int32
	PricePerKG int32
	SlugName   string
	StockKG    float64
}
