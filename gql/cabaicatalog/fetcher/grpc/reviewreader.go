package grpcfetcher

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/data"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/requests"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"strconv"
)

type ReviewReader struct {
	conn connection.CatalogConnection
}

func NewReviewReader(conn connection.CatalogConnection) *ReviewReader {
	return &ReviewReader{conn: conn}
}

func (r *ReviewReader) List(ctx context.Context, req requests.ListReview) ([]*data.Review, error) {
	productID, err := strconv.ParseInt(req.ProductID, 10, 64)
	if err != nil {
		return nil, err
	}
	shopID, err := strconv.ParseInt(req.ShopID, 10, 64)
	if err != nil {
		return nil, err
	}
	res, err := r.conn.ListReviews(ctx, &catalogproto.ListReviewsReq{
		ProductId: productID,
		ShopId:    shopID,
		Pagination: &catalogproto.Pagination{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
	if err != nil {
		return nil, err
	}
	return data.ReviewFromProtos(res.GetReviews()), nil
}
