package mutator

import (
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/request"
)

type CartWriter interface {
	CreateCart(req request.CreateCart) (*data.Cart, error)
	UpdateCart(req request.UpdateCart) (*data.Cart, error)
	DeleteCart(cartID, userID int64) error
}

type OrderWriter interface {
	Checkout(req request.CheckoutReq) ([]*data.Order, error)
	ShipOrder(orderID, shopID int64) (*data.Order, error)
}
