package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/requests"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
)

type TokenFetcher struct {
	conn connection.Connection
}

func NewTokenFetcher(conn connection.Connection) *TokenFetcher {
	return &TokenFetcher{conn: conn}
}

func (tf *TokenFetcher) Login(ctx context.Context, req requests.Login) (*data.Token, error) {
	got, err := tf.conn.Login(ctx, &authproto.LoginReq{
		UsernameOrEmail: req.UserNameOrEmail,
		Password:        req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &data.Token{Token: got.GetToken()}, nil
}

func (tf *TokenFetcher) RefreshToken(ctx context.Context, req data.Token) (*data.Token, error) {
	got, err := tf.conn.RefreshToken(ctx, &authproto.RefreshTokenReq{
		Token: req.Token,
	})
	if err != nil {
		return nil, err
	}
	return &data.Token{Token: got.GetToken()}, nil
}
