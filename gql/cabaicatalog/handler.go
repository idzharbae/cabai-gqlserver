package cabaicatalog

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/fetcher"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/resolver"
)

type CabaiCatalogHandler struct {
	productReader fetcher.ProductReader
}

func NewCabaiCatalogHandler(productReader fetcher.ProductReader) *CabaiCatalogHandler {
	return &CabaiCatalogHandler{productReader: productReader}
}

func (r *CabaiCatalogHandler) Products(ctx context.Context, args struct {
	Params requests.ListProduct
}) (*[]*resolver.Product, error) {
	res := r.productReader.ListProducts(context.Background(), args.Params)
	products := resolver.NewProducts(res)
	return &products, nil
}
