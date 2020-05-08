package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
	"github.com/idzharbae/cabai-gqlserver/util"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

type OrderReader struct {
	conn connection.TransactionConnection
}

func NewOrderReader(conn connection.TransactionConnection) *OrderReader {
	return &OrderReader{conn: conn}
}

func (or *OrderReader) CustomerOrders(customerID int64, status string) ([]*data.Order, error) {
	orderStatusCode := util.OrderStringToCode(status)
	res, err := or.conn.ListOrder(context.Background(), &prototransaction.ListOrderReq{
		CustomerId: customerID,
		Status:     int32(orderStatusCode),
	})
	if err != nil {
		return nil, err
	}
	return data.OrdersFromProtos(res.GetOrders()), nil
}
func (or *OrderReader) ShopOrders(shopID int64, status string) ([]*data.Order, error) {
	orderStatusCode := util.OrderStringToCode(status)
	res, err := or.conn.ListOrder(context.Background(), &prototransaction.ListOrderReq{
		ShopId: shopID,
		Status: int32(orderStatusCode),
	})
	if err != nil {
		return nil, err
	}
	return data.OrdersFromProtos(res.GetOrders()), nil
}
