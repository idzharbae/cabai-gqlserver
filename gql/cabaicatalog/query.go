package cabaicatalog

const (
	Schema = `
		schema {
			query: Query
		}
		# List Cabai Products
		type Query{
			products(params: ProductInput!): [Product]
		}
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
