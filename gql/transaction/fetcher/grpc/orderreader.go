package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/request"
	"github.com/idzharbae/cabai-gqlserver/util"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
)

type OrderReader struct {
	conn connection.TransactionConnection
}

func NewOrderReader(conn connection.TransactionConnection) *OrderReader {
	return &OrderReader{conn: conn}
}

func (or *OrderReader) CustomerOrders(req request.ListOrder) ([]*data.Order, error) {
	orderStatusCode := util.OrderStringToCode(req.Status)
	res, err := or.conn.ListOrder(context.Background(), &prototransaction.ListOrderReq{
		CustomerId: int64(req.UserID),
		Status:     int32(orderStatusCode),
		Pagination: &prototransaction.Pagination{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
	if err != nil {
		return nil, err
	}
	return data.OrdersFromProtos(res.GetOrders()), nil
}
func (or *OrderReader) ShopOrders(req request.ListOrder) ([]*data.Order, error) {
	orderStatusCode := util.OrderStringToCode(req.Status)
	res, err := or.conn.ListOrder(context.Background(), &prototransaction.ListOrderReq{
		ShopId: int64(req.UserID),
		Status: int32(orderStatusCode),
		Pagination: &prototransaction.Pagination{
			Page:  req.Page,
			Limit: req.Limit,
		},
	})
	if err != nil {
		return nil, err
	}
	return data.OrdersFromProtos(res.GetOrders()), nil
}
