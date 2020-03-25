package grpcfetcher

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
)

type ShopReader struct {
	conn connection.Connection
}

func NewShopReader(conn connection.Connection) *ShopReader {
	return &ShopReader{conn: conn}
}

func (s *ShopReader) Get(ctx context.Context, req requests.GetShop) (*data.Shop, error) {
	res, err := s.conn.GetShop(context.Background(), &catalogproto.GetShopReq{
		Id:   req.ID,
		Slug: req.SlugName,
	})
	if err != nil || res == nil {
		return nil, err
	}
	return data.ShopFromProto(res), nil
}

func (s *ShopReader) List(ctx context.Context, req requests.ListShop) ([]*data.Shop, error) {
	res, err := s.conn.ListShops(context.Background(), &catalogproto.ListShopsReq{
		Pagination: &catalogproto.Pagination{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
	if err != nil || res == nil {
		return nil, err
	}
	return data.ShopsFromProtos(res.GetShops()), nil
}
