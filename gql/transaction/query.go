package transaction

const (
	Query = `
		# List Carts
		carts(token: String = ""): [Cart]
	`
	Mutation = ``
	Types    = `
		type Cart{
			id: ID!
			product:  Product
			userID: Int!
			AmountKG: Float!
		}
`
)
