package grpc

import (
	"context"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/connection"
	"github.com/idzharbae/cabai-gqlserver/gql/auth/data"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
)

type SaldoHistoryReader struct {
	conn connection.Connection
}

func NewSaldoHistoryReader(conn connection.Connection) *SaldoHistoryReader {
	return &SaldoHistoryReader{conn: conn}
}

func (ur *UserReader) ListSaldoHistory(ctx context.Context, req *authproto.ListSaldoHistoryReq) ([]*data.SaldoHistory, error) {
	got, err := ur.conn.ListSaldoHistory(ctx, req)
	if err != nil {
		return nil, err
	}
	return data.SaldoHistoriesFromProtos(got.GetSaldoHistories()), nil
}
