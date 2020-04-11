package data

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"strconv"
	"time"
)

type Product struct {
	ID         string
	ShopID     int32
	Name       string
	Quantity   int32
	PricePerKG int32
	SlugName   string
	StockKG    float64
	PhotoURL   string
	CreatedAt  string
	UpdatedAt  string
}

func ProductsFromProtos(protos []*catalogproto.Product) []*Product {
	prods := make([]*Product, len(protos))
	for i, proto := range protos {
		prods[i] = ProductFromProto(proto)
	}

	return prods
}

func ProductFromProto(proto *catalogproto.Product) *Product {
	return &Product{
		ID:         strconv.Itoa(int(proto.GetId())),
		ShopID:     proto.GetShopId(),
		Name:       proto.GetName(),
		Quantity:   proto.GetQuantity(),
		PricePerKG: proto.GetPricePerKg(),
		StockKG:    float64(proto.GetStockKg()),
		SlugName:   proto.GetSlug(),
		PhotoURL:   proto.GetPhotoUrl(),
		CreatedAt:  time.Unix(proto.GetCreatedAt(), 0).Format(time.RFC3339),
		UpdatedAt:  time.Unix(proto.GetUpdatedAt(), 0).Format(time.RFC3339),
	}
}
