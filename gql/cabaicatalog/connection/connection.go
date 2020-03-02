package connection

import "github.com/idzharbae/marketplace-backend/marketplaceproto"

type Connection interface {
	marketplaceproto.MarketplaceClient
}
