package data

import (
	"github.com/idzharbae/marketplace-backend/marketplaceproto"
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
	CreatedAt  string
	UpdatedAt  string
}

func ProductsFromProtos(protos []*marketplaceproto.Product) []*Product {
	prods := make([]*Product, len(protos))
	for i, proto := range protos {
		prods[i] = ProductFromProto(proto)
	}

	return prods
}

func ProductFromProto(proto *marketplaceproto.Product) *Product {
	return &Product{
		ID:         strconv.Itoa(int(proto.GetID())),
		ShopID:     proto.GetShopID(),
		Name:       proto.GetName(),
		Quantity:   proto.GetQuantity(),
		PricePerKG: proto.GetPricePerKG(),
		StockKG:    float64(proto.GetStockKG()),
		SlugName:   proto.GetSlug(),
		CreatedAt:  time.Unix(proto.GetCreatedAt(), 0).Format(time.RFC3339),
		UpdatedAt:  time.Unix(proto.GetUpdatedAt(), 0).Format(time.RFC3339),
	}
}
