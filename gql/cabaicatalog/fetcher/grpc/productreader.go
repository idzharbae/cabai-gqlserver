package grpcfetcher

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
)

type ProductReader struct {
	conn connection.Connection
}

func NewProductReader(conn connection.Connection) *ProductReader {
	return &ProductReader{conn: conn}
}

func (pr *ProductReader) ListProducts(ctx context.Context, req requests.ListProduct) ([]*data.Product, error) {
	res, err := pr.conn.ListProducts(context.Background(), &catalogproto.ListProductsReq{
		Pagination: &catalogproto.Pagination{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
	if err != nil || res == nil {
		return nil, err
	}
	return data.ProductsFromProtos(res.Products), nil
}
