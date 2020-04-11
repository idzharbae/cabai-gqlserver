package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
)

type UserReader struct {
	conn connection.Connection
}

func NewUserReader(conn connection.Connection) *UserReader {
	return &UserReader{conn: conn}
}

func (ur *UserReader) GetUser(ctx context.Context, req *authproto.GetUserReq) (*data.User, error) {
	got, err := ur.conn.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return data.UserFromProto(got), nil
}
