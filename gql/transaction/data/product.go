package data

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
	"strconv"
)

type Product struct {
	ID         string
	ShopID     int64
	Name       string
	AmountKG   float64
	PricePerKG int32
	TotalPrice int64
	PhotoURL   string
	Slug       string
}

func ProductFromProto(proto *prototransaction.Product) *Product {
	return &Product{
		ID:         strconv.FormatInt(proto.GetId(), 10),
		ShopID:     proto.GetShopId(),
		Name:       proto.GetName(),
		AmountKG:   proto.GetAmountKg(),
		PricePerKG: proto.GetPricePerKg(),
		TotalPrice: proto.GetTotalPrice(),
		PhotoURL:   proto.GetPhotoUrl(),
		Slug:       proto.GetSlug(),
	}
}

func ProductsFromProtos(protos []*prototransaction.Product) []*Product {
	products := make([]*Product, len(protos))
	for i, proto := range protos {
		products[i] = ProductFromProto(proto)
	}
	return products
}
