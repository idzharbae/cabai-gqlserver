package connection

import (
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/resources/protoresources"
)

type Connection interface {
	authproto.MarketplaceAuthClient
}

type CatalogConnection interface {
	catalogproto.MarketplaceCatalogClient
}

type ResourcesConnection interface {
	protoresources.MarketplaceResourcesClient
}
