package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
)

type Product struct {
	Data *data.Product
}

func NewProducts(ds []*data.Product) []*Product {
	prods := make([]*Product, 0, len(ds))
	for _, d := range ds {
		if d == nil {
			continue
		}

		prods = append(prods, NewProduct(d))
	}

	return prods
}

func NewProduct(data *data.Product) *Product {
	return &Product{Data: data}
}

func (p *Product) ID() graphql.ID {
	return graphql.ID(p.Data.ID)
}

func (p *Product) ShopID() int32 {
	return p.Data.ShopID
}

func (p *Product) Name() string {
	return p.Data.Name
}

func (p *Product) Quantity() int32 {
	return p.Data.Quantity
}

func (p *Product) PricePerKG() int32 {
	return p.Data.PricePerKG
}

func (p *Product) StockKG() float64 {
	return p.Data.StockKG
}

func (p *Product) CreatedAt() string {
	return p.Data.CreatedAt
}

func (p *Product) UpdatedAt() string {
	return p.Data.UpdatedAt
}

func (p *Product) SlugName() string {
	return p.Data.SlugName
}

func (p *Product) PhotoURL() string {
	return p.Data.PhotoURL
}

func (p *Product) Description() string {
	return p.Data.Description
}

func (p *Product) Category() string {
	return p.Data.Category
}

func (p *Product) BoughtKG() float64 {
	return 0
}

func (p *Product) TotalReviews() int32 {
	return p.Data.TotalReviews
}
func (p *Product) AverageRating() float64 {
	return p.Data.AverageRating
}
