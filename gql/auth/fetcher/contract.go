package fetcher

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/requests"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
)

type TokenFetcher interface {
	Login(ctx context.Context, req requests.Login) (*data.Token, error)
	RefreshToken(ctx context.Context, req data.Token) (*data.Token, error)
}

type UserReader interface {
	GetUser(ctx context.Context, req *authproto.GetUserReq) (*data.User, error)
}

type SaldoHistoryReader interface {
	ListSaldoHistory(ctx context.Context, req *authproto.ListSaldoHistoryReq) ([]*data.SaldoHistory, error)
}
