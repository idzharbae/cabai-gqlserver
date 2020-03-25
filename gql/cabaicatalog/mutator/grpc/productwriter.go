package grpcmutator

import (
	"context"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"

	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
)

type ProductWriter struct {
	conn connection.Connection
}

func NewProductWriter(conn connection.Connection) *ProductWriter {
	return &ProductWriter{conn: conn}
}

func (pr *ProductWriter) CreateProduct(ctx context.Context, req requests.CreateProduct) (*data.Product, error) {
	res, err := pr.conn.CreateProduct(context.Background(), &catalogproto.Product{
		ShopId:     req.ShopID,
		Name:       req.Name,
		Quantity:   req.Quantity,
		PricePerKg: req.PricePerKG,
		StockKg:    float32(req.StockKG),
		Slug:       req.SlugName,
	})
	if err != nil {
		return nil, err
	}
	product := data.ProductFromProto(res)
	return product, nil
}

func (pr *ProductWriter) UpdateProduct(ctx context.Context, req requests.UpdateProduct) (*data.Product, error) {
	res, err := pr.conn.UpdateProduct(context.Background(), &catalogproto.Product{
		Id:         req.ID,
		ShopId:     req.ShopID,
		Name:       req.Name,
		Quantity:   req.Quantity,
		PricePerKg: req.PricePerKG,
		StockKg:    float32(req.StockKG),
		Slug:       req.SlugName,
	})
	if err != nil {
		return nil, err
	}
	product := data.ProductFromProto(res)
	return product, nil
}

func (pr *ProductWriter) DeleteProduct(ctx context.Context, req requests.GetProduct) error {
	_, err := pr.conn.DeleteProduct(context.Background(), &catalogproto.GetProductReq{
		Id:   req.ID,
		Slug: req.SlugName,
	})
	return err
}
