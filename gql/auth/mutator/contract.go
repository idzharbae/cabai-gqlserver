package mutator

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/requests"
)

type UserWriter interface {
	Register(ctx context.Context, user requests.Register) (*data.User, error)
	EditProfile(ctx context.Context, req requests.EditProfile) (*data.User, error)
	TopupSaldo(userID, amount int64) (*data.User, error)
}
