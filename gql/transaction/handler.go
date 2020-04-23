package transaction

import (
	"context"
	authdata "github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/fetcher"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/resolver"
	"github.com/idzharbae/cabai-gqlserver/util"
	"strconv"
)

type TransactionHandler struct {
	cartReader fetcher.CartReader
}

func NewTransactionHandler(cartReader fetcher.CartReader) *TransactionHandler {
	return &TransactionHandler{cartReader: cartReader}
}

func (h *TransactionHandler) Carts(ctx context.Context, args struct {
	Token string
}) (*[]*resolver.Cart, error) {
	var token string
	var err error

	if args.Token == "" {
		token, err = util.GetTokenFromContext(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		token = args.Token
	}

	user, err := authdata.UserFromToken(token)
	if err != nil {
		return nil, err
	}
	userID, err := strconv.ParseInt(user.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	res, err := h.cartReader.ListByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	carts := resolver.NewCarts(res)
	return &carts, nil
}
