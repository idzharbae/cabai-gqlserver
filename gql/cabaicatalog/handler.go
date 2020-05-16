package cabaicatalog

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/fetcher"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/mutator"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/resolver"
)

type CabaiCatalogHandler struct {
	productReader fetcher.ProductReader
	productWriter mutator.ProductWriter
	reviewReader  fetcher.ReviewReader
}

func NewCabaiCatalogHandler(productReader fetcher.ProductReader, productWriter mutator.ProductWriter,
	reviewReader fetcher.ReviewReader) *CabaiCatalogHandler {
	return &CabaiCatalogHandler{
		productReader: productReader,
		productWriter: productWriter,
		reviewReader:  reviewReader,
	}
}

func (r *CabaiCatalogHandler) Reviews(ctx context.Context, args struct {
	Params requests.ListReview
}) (*[]*resolver.Review, error) {
	res, err := r.reviewReader.List(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	reviews := resolver.NewReviews(res)
	return &reviews, nil
}

func (r *CabaiCatalogHandler) SearchProducts(ctx context.Context, args struct {
	Params requests.ListProduct
}) (*[]*resolver.Product, error) {
	res, err := r.productReader.Search(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	products := resolver.NewProducts(res)
	return &products, nil
}

func (r *CabaiCatalogHandler) ProductsByShop(ctx context.Context, args struct {
	Params requests.ProductsByShop
}) (*[]*resolver.Product, error) {
	res, err := r.productReader.GetByShopID(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	products := resolver.NewProducts(res)
	return &products, nil
}

func (r *CabaiCatalogHandler) Product(ctx context.Context, args struct {
	Params requests.GetProduct
}) (*resolver.Product, error) {
	res, err := r.productReader.Get(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	product := resolver.NewProduct(res)
	return product, nil
}

func (r *CabaiCatalogHandler) CreateProduct(ctx context.Context, args struct {
	Params requests.CreateProduct
}) (*resolver.Product, error) {
	res, err := r.productWriter.CreateProduct(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	product := resolver.NewProduct(res)
	return product, nil
}

func (r *CabaiCatalogHandler) UpdateProduct(ctx context.Context, args struct {
	Params requests.UpdateProduct
}) (*resolver.Product, error) {
	res, err := r.productWriter.UpdateProduct(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	product := resolver.NewProduct(res)
	return product, nil
}

func (r *CabaiCatalogHandler) DeleteProduct(ctx context.Context, args struct {
	Params requests.GetProduct
}) (*resolver.Success, error) {
	err := r.productWriter.DeleteProduct(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	return &resolver.Success{}, nil
}
