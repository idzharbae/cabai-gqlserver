package auth

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/util"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"strconv"

	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/fetcher"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/mutator"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/requests"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/resolver"
)

type AuthHandler struct {
	tokenFetcher fetcher.TokenFetcher
	userWriter   mutator.UserWriter
	userReader   fetcher.UserReader
}

func NewAuthHandler(tokenFetcher fetcher.TokenFetcher, userWriter mutator.UserWriter, userReader fetcher.UserReader) *AuthHandler {
	return &AuthHandler{
		tokenFetcher: tokenFetcher,
		userWriter:   userWriter,
		userReader:   userReader,
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

func (ah *AuthHandler) GetUserInfo(ctx context.Context, args struct {
	Token *string
}) (*resolver.User, error) {
	var token string
	var err error

	if args.Token != nil {
		token = *args.Token
	} else {
		token, err = util.GetTokenFromContext(ctx)
		if err != nil {
			return nil, err
		}
	}

	userFromToken, err := data.UserFromToken(token)
	if err != nil {
		return nil, err
	}
	id, err := strconv.ParseInt(userFromToken.ID, 10, 64)
	user, err := ah.userReader.GetUser(ctx, &authproto.GetUserReq{
		Id: id,
	})
	if err != nil || user == nil {
		return nil, err
	}
	userFromToken.PhotoURL = user.PhotoURL
	return resolver.NewUser(&userFromToken), nil
}
