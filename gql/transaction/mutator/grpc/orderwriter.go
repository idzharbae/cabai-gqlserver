package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/request"
	"github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"
	"strconv"
)

type OrderWriter struct {
	conn connection.TransactionConnection
}

func NewOrderWriter(conn connection.TransactionConnection) *OrderWriter {
	return &OrderWriter{conn: conn}
}

func (ow *OrderWriter) Checkout(req request.CheckoutReq) ([]*data.Order, error) {
	userID, err := strconv.ParseInt(req.UserID, 10, 64)
	if err != nil {
		return nil, err
	}
	paymentAmount, err := strconv.ParseInt(req.PaymentAmount, 10, 64)
	if err != nil {
		return nil, err
	}
	cartIDs, err := stringSliceToIntSlice(req.CartIDs)
	if err != nil {
		return nil, err
	}
	res, err := ow.conn.Checkout(context.Background(), &prototransaction.CheckoutReq{
		UserId:        userID,
		CartIds:       cartIDs,
		PaymentAmount: paymentAmount,
	})
	if err != nil {
		return nil, err
	}
	return data.OrdersFromProtos(res.GetOrders()), nil
}

func (ow *OrderWriter) ShipOrder(orderID, shopID int64) (*data.Order, error) {
	res, err := ow.conn.UpdateOrderStatusToOnShipment(context.Background(), &prototransaction.ChangeProductStatusReq{
		OrderId: orderID,
		ShopId:  shopID,
	})
	if err != nil {
		return nil, err
	}
	return data.OrderFromProto(res.GetOrder()), nil
}
func (ow *OrderWriter) RejectOrder(orderID, shopID int64) (*data.Order, error) {
	res, err := ow.conn.RejectOrder(context.Background(), &prototransaction.ChangeProductStatusReq{
		OrderId: orderID,
		ShopId:  shopID,
	})
	if err != nil {
		return nil, err
	}
	return data.OrderFromProto(res), nil
}
func (ow *OrderWriter) FulfillOrder(orderID, userID int64) error {
	_, err := ow.conn.Fulfill(context.Background(), &prototransaction.FulfillReq{
		OrderId: orderID,
		UserId:  userID,
	})
	return err
}

func stringSliceToIntSlice(s []string) ([]int64, error) {
	res := make([]int64, len(s))
	for i, str := range s {
		integer, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}
		res[i] = integer
	}
	return res, nil
}
