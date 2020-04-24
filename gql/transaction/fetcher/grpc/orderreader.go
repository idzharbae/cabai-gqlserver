package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

type OrderReader struct {
	conn connection.TransactionConnection
}

func NewOrderReader(conn connection.TransactionConnection) *OrderReader {
	return &OrderReader{conn: conn}
}

func (or *OrderReader) CustomerOrders(customerID int64) ([]*data.Order, error) {
	res, err := or.conn.ListOrder(context.Background(), &prototransaction.ListOrderReq{
		CustomerId: customerID,
	})
	if err != nil {
		return nil, err
	}
	return data.OrdersFromProtos(res.GetOrders()), nil
}
func (or *OrderReader) ShopOrders(shopID int64) ([]*data.Order, error) {
	res, err := or.conn.ListOrder(context.Background(), &prototransaction.ListOrderReq{
		ShopId: shopID,
	})
	if err != nil {
		return nil, err
	}
	return data.OrdersFromProtos(res.GetOrders()), nil
}
