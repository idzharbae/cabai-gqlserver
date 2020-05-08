package transaction

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/fetcher"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/mutator"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/request"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/resolver"
	"github.com/idzharbae/cabai-gqlserver/util"
	"strconv"
)

type TransactionHandler struct {
	cartReader  fetcher.CartReader
	orderReader fetcher.OrderReader
	cartWriter  mutator.CartWriter
	orderWriter mutator.OrderWriter
}

func NewTransactionHandler(cartReader fetcher.CartReader, cartWriter mutator.CartWriter,
	orderWriter mutator.OrderWriter, orderReader fetcher.OrderReader) *TransactionHandler {
	return &TransactionHandler{cartReader: cartReader, cartWriter: cartWriter, orderWriter: orderWriter, orderReader: orderReader}
}

func (h *TransactionHandler) Carts(ctx context.Context, args struct {
	Token string
}) (*[]*resolver.Cart, error) {
	userID, err := h.getUserID(args.Token, ctx)
	if err != nil {
		return nil, err
	}

	res, err := h.cartReader.ListByUserID(userID)
	if err != nil {
		return nil, err
	}
	carts := resolver.NewCarts(res)
	return &carts, nil
}

func (h *TransactionHandler) CreateCart(ctx context.Context, args struct {
	Params request.CreateCart
}) (*resolver.Cart, error) {
	userID, err := h.getUserID("", ctx)
	if err != nil {
		return nil, err
	}
	cart, err := h.cartWriter.CreateCart(request.CreateCart{
		ProductId:  args.Params.ProductId,
		UserId:     userID,
		QuantityKg: args.Params.QuantityKg,
	})
	if err != nil {
		return nil, err
	}
	return resolver.NewCart(cart), nil
}

func (h *TransactionHandler) UpdateCartQuantity(ctx context.Context, args struct {
	Params request.UpdateCart
}) (*resolver.Cart, error) {
	userID, err := h.getUserID("", ctx)
	if err != nil {
		return nil, err
	}
	cart, err := h.cartWriter.UpdateCart(request.UpdateCart{
		CartID:        args.Params.CartID,
		UserId:        userID,
		NewQuantityKG: args.Params.NewQuantityKG,
	})
	if err != nil {
		return nil, err
	}
	return resolver.NewCart(cart), nil
}

func (h *TransactionHandler) DeleteCart(ctx context.Context, args struct {
	CartID int32
}) (*resolver.Success, error) {
	userID, err := h.getUserID("", ctx)
	if err != nil {
		return nil, err
	}
	err = h.cartWriter.DeleteCart(int64(args.CartID), userID)
	if err != nil {
		return nil, err
	}
	return &resolver.Success{}, nil
}

func (h *TransactionHandler) Checkout(ctx context.Context, args struct {
	Params request.CheckoutReq
}) (*[]*resolver.Order, error) {
	userID, err := h.getUserID("", ctx)
	if err != nil {
		return nil, err
	}
	args.Params.UserID = strconv.FormatInt(userID, 10)
	res, err := h.orderWriter.Checkout(args.Params)
	if err != nil {
		return nil, err
	}
	orders := resolver.NewOrders(res)
	return &orders, nil
}
func (h *TransactionHandler) CustomerOrders(ctx context.Context, args struct {
	Status string
}) (*[]*resolver.Order, error) {
	userID, err := h.getUserID("", ctx)
	if err != nil {
		return nil, err
	}

	res, err := h.orderReader.CustomerOrders(userID, args.Status)
	if err != nil {
		return nil, err
	}
	orders := resolver.NewOrders(res)
	return &orders, nil
}
func (h *TransactionHandler) ShopOrders(ctx context.Context, args struct {
	Status string
}) (*[]*resolver.Order, error) {
	userID, err := h.getUserID("", ctx)
	if err != nil {
		return nil, err
	}

	res, err := h.orderReader.ShopOrders(userID, args.Status)
	if err != nil {
		return nil, err
	}
	orders := resolver.NewOrders(res)
	return &orders, nil
}
func (h *TransactionHandler) ShipOrder(ctx context.Context, args struct {
	OrderID int32
}) (*resolver.Order, error) {
	userID, err := h.getUserID("", ctx)
	if err != nil {
		return nil, err
	}

	res, err := h.orderWriter.ShipOrder(int64(args.OrderID), userID)
	if err != nil {
		return nil, err
	}
	order := resolver.NewOrder(res)
	return order, nil
}
func (h *TransactionHandler) RejectOrder(ctx context.Context, args struct {
	OrderID int32
}) (*resolver.Order, error) {
	userID, err := h.getUserID("", ctx)
	if err != nil {
		return nil, err
	}

	res, err := h.orderWriter.RejectOrder(int64(args.OrderID), userID)
	if err != nil {
		return nil, err
	}
	order := resolver.NewOrder(res)
	return order, nil
}
func (h *TransactionHandler) FulfillOrder(ctx context.Context, args struct {
	OrderID int32
}) (*resolver.Success, error) {
	userID, err := h.getUserID("", ctx)
	if err != nil {
		return nil, err
	}

	err = h.orderWriter.FulfillOrder(int64(args.OrderID), userID)
	if err != nil {
		return nil, err
	}
	return &resolver.Success{}, nil
}

func (h *TransactionHandler) getUserID(token string, ctx context.Context) (int64, error) {
	var userID int64
	var err error

	if token != "" {
		userID, err = util.UserIDFromToken(token)
	} else {
		userID, err = util.UserIDFromCtx(ctx)
	}
	if err != nil {
		return 0, err
	}
	return userID, nil
}
