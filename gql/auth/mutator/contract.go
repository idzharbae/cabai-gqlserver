package mutator

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/requests"
)

type UserWriter interface {
	Register(ctx context.Context, user requests.Register) (*data.User, error)
}
