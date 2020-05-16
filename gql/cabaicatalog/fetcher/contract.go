package fetcher

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
)

type ProductReader interface {
	Search(ctx context.Context, req requests.ListProduct) ([]*data.Product, error)
	GetByShopID(ctx context.Context, req requests.ProductsByShop) ([]*data.Product, error)
	Get(ctx context.Context, req requests.GetProduct) (*data.Product, error)
}

type ReviewReader interface {
	List(ctx context.Context, req requests.ListReview) ([]*data.Review, error)
}
