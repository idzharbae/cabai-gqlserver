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
	shopReader    fetcher.ShopReader
}

func NewCabaiCatalogHandler(productReader fetcher.ProductReader, productWriter mutator.ProductWriter, shopReader fetcher.ShopReader) *CabaiCatalogHandler {
	return &CabaiCatalogHandler{
		productReader: productReader,
		productWriter: productWriter,
		shopReader:    shopReader,
	}
}

func (r *CabaiCatalogHandler) Products(ctx context.Context, args struct {
	Params requests.ListProduct
}) (*[]*resolver.Product, error) {
	res, err := r.productReader.List(ctx, args.Params)
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

func (r *CabaiCatalogHandler) Shop(ctx context.Context, args struct {
	Params requests.GetShop
}) (*resolver.Shop, error) {
	got, err := r.shopReader.Get(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	return resolver.NewShop(got), nil
}

func (r *CabaiCatalogHandler) Shops(ctx context.Context, args struct {
	Params requests.ListShop
}) (*[]*resolver.Shop, error) {
	got, err := r.shopReader.List(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	shops := resolver.NewShops(got)
	return &shops, nil
}
