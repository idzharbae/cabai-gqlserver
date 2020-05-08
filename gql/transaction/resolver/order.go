package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
	"github.com/idzharbae/cabai-gqlserver/util"
)

type Order struct {
	Data *data.Order
}

func NewOrders(ds []*data.Order) []*Order {
	orders := make([]*Order, 0, len(ds))
	for _, d := range ds {
		if d == nil {
			continue
		}

		orders = append(orders, NewOrder(d))
	}

	return orders
}
func NewOrder(order *data.Order) *Order {
	return &Order{Data: order}
}

func (o *Order) ID() graphql.ID {
	return graphql.ID(o.Data.ID)
}
func (o *Order) CustomerID() graphql.ID {
	return graphql.ID(o.Data.UserID)
}
func (o *Order) ShopID() graphql.ID {
	return graphql.ID(o.Data.ShopID)
}
func (o *Order) TotalPrice() string {
	return o.Data.TotalPrice
}
func (o *Order) Status() string {
	return util.OrderCodeToString(int(o.Data.Status))
}
func (o *Order) Products() *[]*Product {
	products := NewProducts(o.Data.Products)
	return &products
}
func (o *Order) Payment() *Payment {
	return NewPayment(o.Data.Payment)
}
