package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/requests"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
)

type UserMutator struct {
	conn connection.Connection
}

func NewUserMutator(conn connection.Connection) *UserMutator {
	return &UserMutator{conn: conn}
}

func (um *UserMutator) Register(ctx context.Context, req requests.Register) error {
	_, err := um.conn.Register(ctx, &authproto.RegisterReq{
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.PhoneNumber,
		Password: req.Password,
		Type:     req.Role,
		FullName: req.FullName,
	})
	return err
}
