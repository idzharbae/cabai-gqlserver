package connection

import "github.com/idzharbae/marketplace-backend/svc/transaction/prototransaction"

type TransactionConnection interface {
	prototransaction.MarketplaceTransactionClient
}
