package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
)

type Cart struct {
	Data *data.Cart
}

func NewCarts(ds []*data.Cart) []*Cart {
	carts := make([]*Cart, 0, len(ds))
	for _, d := range ds {
		if d == nil {
			continue
		}

		carts = append(carts, NewCart(d))
	}

	return carts
}

func NewCart(data *data.Cart) *Cart {
	return &Cart{Data: data}
}

func (c *Cart) ID() graphql.ID {
	return graphql.ID(c.Data.ID)
}
func (c *Cart) Product() *Product {
	return NewProduct(c.Data.Product)
}
func (c *Cart) UserID() int32 {
	return int32(c.Data.UserID)
}
func (c *Cart) AmountKG() float64 {
	return c.Data.AmountKG
}
