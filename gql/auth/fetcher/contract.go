package fetcher

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/requests"
)

type TokenFetcher interface {
	Login(ctx context.Context, req requests.Login) (*data.Token, error)
	RefreshToken(ctx context.Context, req data.Token) (*data.Token, error)
}
