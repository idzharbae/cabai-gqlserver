package data

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
	"strconv"
)

type Cart struct {
	ID       string
	Product  *Product
	UserID   int64
	AmountKG float64
}

func CartsFromProtos(protos []*prototransaction.Cart) []*Cart {
	carts := make([]*Cart, len(protos))
	for i, proto := range protos {
		carts[i] = CartFromProto(proto)
	}
	return carts
}

func CartFromProto(proto *prototransaction.Cart) *Cart {
	return &Cart{
		ID:       strconv.FormatInt(proto.GetId(), 10),
		Product:  ProductFromProto(proto.GetProduct()),
		UserID:   proto.GetUserId(),
		AmountKG: proto.GetProduct().GetAmountKg(),
	}
}
