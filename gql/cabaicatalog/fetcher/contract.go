package fetcher

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
)

type ProductReader interface {
	ListProducts(ctx context.Context, req requests.ListProduct) ([]*data.Product, error)
}
