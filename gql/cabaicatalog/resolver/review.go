package resolver

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
)

type Review struct {
	Data *data.Review
}

func NewReviews(ds []*data.Review) []*Review {
	prods := make([]*Review, 0, len(ds))
	for _, d := range ds {
		if d == nil {
			continue
		}

		prods = append(prods, NewReview(d))
	}

	return prods
}

func NewReview(data *data.Review) *Review {
	return &Review{Data: data}
}

func (r *Review) ID() graphql.ID {
	return graphql.ID(r.Data.ID)
}
func (r *Review) ProductID() graphql.ID {
	return graphql.ID(r.Data.ProductID)
}
func (r *Review) UserID() graphql.ID {
	return graphql.ID(r.Data.UserID)
}
func (r *Review) ShopID() graphql.ID {
	return graphql.ID(r.Data.ShopID)
}
func (r *Review) Title() string {
	return r.Data.Title
}
func (r *Review) Content() string {
	return r.Data.Content
}
func (r *Review) PhotoURL() string {
	return r.Data.PhotoURL
}
func (r *Review) Rating() float64 {
	return r.Data.Rating
}
func (r *Review) CreatedAt() string {
	return r.Data.CreatedAt
}
func (r *Review) UpdatedAt() string {
	return r.Data.UpdatedAt
}
