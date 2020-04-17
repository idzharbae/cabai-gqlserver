package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/requests"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
)

type UserMutator struct {
	conn    connection.Connection
	catalog connection.CatalogConnection
}

func NewUserMutator(conn connection.Connection, catalog connection.CatalogConnection) *UserMutator {
	return &UserMutator{conn: conn, catalog: catalog}
}

func (um *UserMutator) Register(ctx context.Context, req requests.Register) (*data.User, error) {
	user, err := um.conn.Register(ctx, &authproto.RegisterReq{
		UserName:      req.UserName,
		Email:         req.Email,
		Phone:         req.PhoneNumber,
		Password:      req.Password,
		Type:          req.Role,
		FullName:      req.FullName,
		City:          req.City,
		Province:      req.Province,
		AddressDetail: req.AddressDetail,
		PhotoUrl:      req.PhotoURL,
		ZipCode:       req.ZipCode,
	})
	if err != nil {
		return nil, err
	}
	userData := data.UserFromProto(user)
	return userData, nil
}
