package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/fetcher/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/marketplace-backend/marketplaceproto"
)

type ProductReader struct {
	conn connection.Connection
}

func NewProductReader(conn connection.Connection) *ProductReader {
	return &ProductReader{conn: conn}
}

func (pr *ProductReader) ListProducts(ctx context.Context, req requests.ListProduct) []*data.Product {
	res, err := pr.conn.ListProducts(context.Background(), &marketplaceproto.ListProductsReq{
		Pagination: &marketplaceproto.Pagination{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil || res == nil {
		return nil
	}
	return data.ProductsFromProtos(res.Products)
}
