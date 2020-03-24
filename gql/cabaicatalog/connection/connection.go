package connection

import "github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"

type Connection interface {
	catalogproto.MarketplaceCatalogClient
}
