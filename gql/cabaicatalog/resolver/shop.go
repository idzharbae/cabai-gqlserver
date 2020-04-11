package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
)

type Shop struct {
	Data *data.Shop
}

func NewShops(ds []*data.Shop) []*Shop {
	var prods []*Shop
	for _, d := range ds {
		if d == nil {
			continue
		}

		prods = append(prods, NewShop(d))
	}

	return prods
}

func NewShop(data *data.Shop) *Shop {
	return &Shop{Data: data}
}

func (s *Shop) ID() graphql.ID {
	return graphql.ID(s.Data.ID)
}
func (s *Shop) Name() string {
	return s.Data.Name
}
func (s *Shop) Address() string {
	return s.Data.Address
}
func (s *Shop) Slug() string {
	return s.Data.SlugName
}
func (s *Shop) Location() *Location {
	return NewLocation(&s.Data.Location)
}
func (s *Shop) Products() *[]*Product {
	prods := NewProducts(s.Data.Products)
	return &prods
}
func (s *Shop) CreatedAt() string {
	return s.Data.CreatedAt
}
func (s *Shop) UpdatedAt() string {
	return s.Data.UpdatedAt
}

func (s *Shop) PhotoURL() string {
	return s.Data.PhotoURL
}
