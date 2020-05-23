package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
	"strconv"
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
	return int32(p.Data.ShopID)
}
func (p *Product) Name() string {
	return p.Data.Name
}
func (p *Product) AmountKG() float64 {
	return p.Data.AmountKG
}
func (p *Product) PricePerKG() int32 {
	return p.Data.PricePerKG
}
func (p *Product) TotalPrice() string {
	return strconv.FormatInt(p.Data.TotalPrice, 10)
}
func (p *Product) PhotoURL() string {
	return p.Data.PhotoURL
}
func (p *Product) BoughtKG() float64 {
	return p.Data.AmountKG
}
func (p *Product) Quantity() int32 {
	return 0
}
func (p *Product) StockKG() float64 {
	return 0
}
func (p *Product) SlugName() string {
	return p.Data.Slug
}
func (p *Product) CreatedAt() string {
	return ""
}
func (p *Product) UpdatedAt() string {
	return ""
}
func (p *Product) Description() string {
	return ""
}
func (p *Product) Category() string {
	return ""
}
func (p *Product) TotalReviews() int32 {
	return 0
}
func (p *Product) AverageRating() float64 {
	return 0
}
