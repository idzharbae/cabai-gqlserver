package transaction

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/fetcher"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/mutator"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/request"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/resolver"
	"github.com/idzharbae/cabai-gqlserver/util"
)

type TransactionHandler struct {
	cartReader fetcher.CartReader
	cartWriter mutator.CartWriter
}

func NewTransactionHandler(cartReader fetcher.CartReader, cartWriter mutator.CartWriter) *TransactionHandler {
	return &TransactionHandler{cartReader: cartReader, cartWriter: cartWriter}
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
