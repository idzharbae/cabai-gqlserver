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
	tokenFetcher       fetcher.TokenFetcher
	userWriter         mutator.UserWriter
	userReader         fetcher.UserReader
	saldoHistoryReader fetcher.SaldoHistoryReader
}

func NewAuthHandler(tokenFetcher fetcher.TokenFetcher, userWriter mutator.UserWriter, userReader fetcher.UserReader,
	saldoHistoryReader fetcher.SaldoHistoryReader) *AuthHandler {
	return &AuthHandler{
		tokenFetcher:       tokenFetcher,
		userWriter:         userWriter,
		userReader:         userReader,
		saldoHistoryReader: saldoHistoryReader,
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
}) (*resolver.User, error) {
	user, err := ah.userWriter.Register(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	return resolver.NewUser(user), nil
}

func (ah *AuthHandler) EditProfile(ctx context.Context, args struct {
	Params requests.EditProfile
}) (*resolver.User, error) {
	user, err := ah.userWriter.EditProfile(ctx, args.Params)
	if err != nil {
		return nil, err
	}
	return resolver.NewUser(user), nil
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
	return resolver.NewUser(user), nil
}

func (ah *AuthHandler) GetUserByID(ctx context.Context, args struct {
	UserID int32
}) (*resolver.User, error) {
	user, err := ah.userReader.GetUser(ctx, &authproto.GetUserReq{
		Id: int64(args.UserID),
	})
	if err != nil || user == nil {
		return nil, err
	}
	return resolver.NewUser(user), nil
}
func (ah *AuthHandler) Topup(ctx context.Context, args struct {
	Amount string
}) (*resolver.User, error) {
	userID, err := util.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	amount, err := strconv.ParseInt(args.Amount, 10, 64)
	if err != nil {
		return nil, err
	}
	user, err := ah.userWriter.TopupSaldo(userID, amount)
	if err != nil || user == nil {
		return nil, err
	}
	return resolver.NewUser(user), nil
}
func (ah *AuthHandler) SaldoHistory(ctx context.Context, args struct {
	Params requests.SaldoHistory
}) (*[]*resolver.SaldoHistory, error) {
	userID, err := util.UserIDFromCtx(ctx)
	if err != nil {
		return nil, err
	}
	histories, err := ah.saldoHistoryReader.ListSaldoHistory(ctx, &authproto.ListSaldoHistoryReq{
		UserId: userID,
		Pagination: &authproto.Pagination{
			Page:  args.Params.Page,
			Limit: args.Params.Limit,
		},
	})
	if err != nil {
		return nil, err
	}
	saldoHistories := resolver.NewSaldoHistories(histories)
	return &saldoHistories, nil
}
