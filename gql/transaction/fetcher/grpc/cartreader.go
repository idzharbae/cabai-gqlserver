package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

type CartReader struct {
	conn connection.TransactionConnection
}

func NewCartReader(conn connection.TransactionConnection) *CartReader {
	return &CartReader{conn: conn}
}

func (cr *CartReader) ListByUserID(userID int64) ([]*data.Cart, error) {
	res, err := cr.conn.ListCartItems(context.Background(), &prototransaction.ListCartItemsReq{
		UserId: userID,
	})
	if err != nil {
		return nil, err
	}
	return data.CartsFromProtos(res.GetCart()), nil
}
