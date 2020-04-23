package fetcher

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/transaction/data"
)

type CartReader interface {
	ListByUserID(ctx context.Context, userID int64) ([]*data.Cart, error)
}
