package auth

import (
	"context"

	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/fetcher"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/mutator"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/requests"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/resolver"
)

type AuthHandler struct {
	tokenFetcher fetcher.TokenFetcher
	userWriter   mutator.UserWriter
}

func NewAuthHandler(tokenFetcher fetcher.TokenFetcher, userWriter mutator.UserWriter) *AuthHandler {
	return &AuthHandler{
		tokenFetcher: tokenFetcher,
		userWriter:   userWriter,
	}
}

func (ah *AuthHandler) Login(ctx context.Context, args struct {
	Params requests.Login
}) (*resolver.Token, error) {
	res, err := ah.tokenFetcher.Login(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	return resolver.NewToken(res), nil
}

func (ah *AuthHandler) RefreshToken(ctx context.Context, args struct {
	Params data.Token
}) (*resolver.Token, error) {
	res, err := ah.tokenFetcher.RefreshToken(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	return resolver.NewToken(res), nil
}

func (ah *AuthHandler) Register(ctx context.Context, args struct {
	Params requests.Register
}) (*resolver.Success, error) {
	err := ah.userWriter.Register(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	return &resolver.Success{}, nil
}
