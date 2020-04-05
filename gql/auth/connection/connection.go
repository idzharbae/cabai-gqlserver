package connection

import "github.com/idzharbae/marketplace-backend/svc/auth/authproto"

type Connection interface {
	authproto.MarketplaceAuthClient
}
