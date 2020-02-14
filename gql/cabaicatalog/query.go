package cabaicatalog

const (
	Schema = `
		# List Cabai Products
		products(params: ProductInput!): [Product]
	`

	Types = `
		type Product {
			id: ID!
			shopID: Int!
			name: String!
			quantity: Int!
			pricePerKG: Int!
			stockKG: Float!
			createdAt: String!
			updatedAt: String!
		}
		input ProductInput{
			page: Int = 1
			limit: Int = 10
		}
`
)
