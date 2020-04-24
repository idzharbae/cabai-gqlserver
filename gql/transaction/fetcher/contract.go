package fetcher

import (
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
)

type CartReader interface {
	ListByUserID(userID int64) ([]*data.Cart, error)
}

type OrderReader interface {
	CustomerOrders(customerID int64) ([]*data.Order, error)
	ShopOrders(shopID int64) ([]*data.Order, error)
}
