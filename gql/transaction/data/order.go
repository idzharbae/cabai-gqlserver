package data

import (
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
	"strconv"
	"time"
)

type Order struct {
	ID         string
	UserID     string
	ShopID     string
	Products   []*Product
	TotalPrice string
	Payment    *Payment
	Status     int32
}

func OrdersFromProtos(protos []*prototransaction.Order) []*Order {
	res := make([]*Order, 0, len(protos))
	for _, proto := range protos {
		if proto == nil {
			continue
		}
		res = append(res, OrderFromProto(proto))
	}
	return res
}

func OrderFromProto(proto *prototransaction.Order) *Order {
	products := ProductsFromProtos(proto.GetProducts())
	return &Order{
		ID:         strconv.FormatInt(proto.GetId(), 10),
		UserID:     strconv.FormatInt(proto.GetUserId(), 10),
		ShopID:     strconv.FormatInt(proto.GetShopId(), 10),
		Products:   products,
		TotalPrice: strconv.FormatInt(proto.GetTotalPrice(), 10),
		Payment:    PaymentFromProto(proto.GetPayment()),
		Status:     proto.GetStatus(),
	}
}

type Payment struct {
	ID            string
	Amount        int64
	PaymentMethod int32
	PaymentStatus int32
	CreatedAt     string
	UpdatedAt     string
}

func PaymentFromProto(proto *prototransaction.Payment) *Payment {
	return &Payment{
		ID:            strconv.FormatInt(proto.GetId(), 10),
		Amount:        proto.GetAmount(),
		PaymentMethod: proto.GetPaymentMethod(),
		PaymentStatus: proto.GetPaymentMethod(),
		CreatedAt:     time.Unix(proto.GetCreatedAt(), 0).Format(time.RFC3339),
		UpdatedAt:     time.Unix(proto.GetUpdatedAt(), 0).Format(time.RFC3339),
	}
}
