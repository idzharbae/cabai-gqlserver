package connection

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
)

type Connection interface {
	authproto.MarketplaceAuthClient
}

type CatalogConnection interface {
	catalogproto.MarketplaceCatalogClient
}
