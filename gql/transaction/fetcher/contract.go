package fetcher

import (
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/request"
)

type CartReader interface {
	ListByUserID(req request.ListCarts) ([]*data.Cart, error)
}

type OrderReader interface {
	CustomerOrders(req request.ListOrder) ([]*data.Order, error)
	ShopOrders(req request.ListOrder) ([]*data.Order, error)
}
