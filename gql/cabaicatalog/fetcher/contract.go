package fetcher

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
)

type ProductReader interface {
	List(ctx context.Context, req requests.ListProduct) ([]*data.Product, error)
	Get(ctx context.Context, req requests.GetProduct) (*data.Product, error)
}

type ShopReader interface {
	List(ctx context.Context, req requests.ListShop) ([]*data.Shop, error)
	Get(ctx context.Context, req requests.GetShop) (*data.Shop, error)
}
