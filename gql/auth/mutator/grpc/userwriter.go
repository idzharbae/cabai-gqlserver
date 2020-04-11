package grpc

import (
	"context"
	"errors"
	"github.com/idzharbae/cabai-gqlserver/globalconstant"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/requests"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
)

type UserMutator struct {
	conn    connection.Connection
	catalog connection.CatalogConnection
}

func NewUserMutator(conn connection.Connection, catalog connection.CatalogConnection) *UserMutator {
	return &UserMutator{conn: conn, catalog: catalog}
}

func (um *UserMutator) Register(ctx context.Context, req requests.Register) error {
	if req.Role == globalconstant.ShopType && req.Address == nil {
		return errors.New("address field is required for shops")
	}

	user, err := um.conn.Register(ctx, &authproto.RegisterReq{
		UserName: req.UserName,
		Email:    req.Email,
		Phone:    req.PhoneNumber,
		Password: req.Password,
		Type:     req.Role,
		FullName: req.FullName,
	})
	if err != nil {
		return err
	}
	if req.Role == globalconstant.ShopType {
		_, err := um.catalog.CreateShop(ctx, &catalogproto.Shop{
			Id:       int32(user.GetId()),
			Name:     user.GetName(),
			Address:  *req.Address,
			Slug:     user.GetUserName(),
			Location: nil,
			Products: nil,
			PhotoUrl: user.GetPhotoUrl(),
		})
		return err
	}
	return err
}
