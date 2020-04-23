package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/request"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

type CartWriter struct {
	conn connection.TransactionConnection
}

func NewCartWriter(conn connection.TransactionConnection) *CartWriter {
	return &CartWriter{conn: conn}
}

func (cw *CartWriter) CreateCart(req request.CreateCart) (*data.Cart, error) {
	res, err := cw.conn.AddToCart(context.Background(), &prototransaction.AddToCartReq{
		ProductId:  int64(req.ProductId),
		UserId:     req.UserId,
		QuantityKg: req.QuantityKg,
	})
	if err != nil {
		return nil, err
	}
	return data.CartFromProto(res), nil
}

func (cw *CartWriter) UpdateCart(req request.UpdateCart) (*data.Cart, error) {
	res, err := cw.conn.UpdateCart(context.Background(), &prototransaction.UpdateCartReq{
		Id:         int64(req.CartID),
		UserId:     req.UserId,
		QuantityKg: req.NewQuantityKG,
	})
	if err != nil {
		return nil, err
	}
	return data.CartFromProto(res), nil
}

func (cw *CartWriter) DeleteCart(cartID, userID int64) error {
	_, err := cw.conn.RemoveCart(context.Background(), &prototransaction.RemoveCartReq{
		Id:     cartID,
		UserId: userID,
	})
	return err
}
