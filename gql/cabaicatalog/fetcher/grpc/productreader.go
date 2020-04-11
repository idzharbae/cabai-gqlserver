package grpcfetcher

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
)

type ProductReader struct {
	conn connection.CatalogConnection
}

func NewProductReader(conn connection.CatalogConnection) *ProductReader {
	return &ProductReader{conn: conn}
}

func (pr *ProductReader) List(ctx context.Context, req requests.ListProduct) ([]*data.Product, error) {
	res, err := pr.conn.ListProducts(context.Background(), &catalogproto.ListProductsReq{
		ShopID: req.ShopID,
		Pagination: &catalogproto.Pagination{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
	if err != nil || res == nil {
		return nil, err
	}
	return data.ProductsFromProtos(res.GetProducts()), nil
}

func (pr *ProductReader) Get(ctx context.Context, req requests.GetProduct) (*data.Product, error) {
	res, err := pr.conn.GetProduct(context.Background(), &catalogproto.GetProductReq{
		Id:   req.ID,
		Slug: req.SlugName,
	})
	if err != nil || res == nil {
		return nil, err
	}
	return data.ProductFromProto(res), nil
}
