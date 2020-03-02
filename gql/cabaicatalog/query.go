package cabaicatalog

const (
	Query = `
		# List Cabai Products
		products(params: ListProductInput!): [Product]
	`
	Mutation = `
	# Create new cabai product
	createProduct(params: CreateProductInput!): Product
	# update cabai product
	updateProduct(params: UpdateProductInput!): Product
	# delete cabai product
	deleteProduct(id: Int!): Success
`
	Types = `
		type Success{
			success: Boolean!
		}
		type Product {
			id: ID!
			shopID: Int!
			name: String!
			quantity: Int!
			pricePerKG: Int!
			stockKG: Float!
			slugName: String!
			createdAt: String!
			updatedAt: String!
		}
		input ListProductInput{
			page: Int = 1
			limit: Int = 10
		}
		input CreateProductInput {
			shopID: Int!
			name: String!
			quantity: Int!
			pricePerKG: Int!
			stockKG: Float!
			slugName: String!
		}
		input UpdateProductInput {
			id: Int!
			shopID: Int!
			name: String!
			quantity: Int!
			pricePerKG: Int!
			stockKG: Float!
			slugName: String!
		}
`
)
