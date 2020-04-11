package connection

import (
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/resources/protoresources"
)

type CatalogConnection interface {
	catalogproto.MarketplaceCatalogClient
}

type ResourcesConnection interface {
	protoresources.MarketplaceResourcesClient
}
