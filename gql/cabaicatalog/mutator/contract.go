package mutator

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
)

type ProductWriter interface {
	CreateProduct(ctx context.Context, req requests.CreateProduct) (*data.Product, error)
	UpdateProduct(ctx context.Context, req requests.UpdateProduct) (*data.Product, error)
	DeleteProduct(ctx context.Context, req int32) error
}
