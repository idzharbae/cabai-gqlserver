package cabaicatalog

const (
	Query = `
		# List Cabai Products
		products(params: ListProductInput!): [Product]
		product(params: GetProductInput!): Product
		shops(params: ListShopInput!): [Shop]
		shop(params: GetShopInput!): Shop
	`
	Mutation = `
	# Create new cabai product
	createProduct(params: CreateProductInput!): Product
	# update cabai product
	updateProduct(params: UpdateProductInput!): Product
	# delete cabai product
	deleteProduct(params: GetProductInput!): Success
`
	Types = `
		scalar Upload
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
			photoURL: String!
			createdAt: String!
			updatedAt: String!
		}
		type Shop {
			id: ID!
			name: String!
			address: String!
			slug: String!
			location: Location!
			products: [Product]
			photoURL: String!
			createdAt: String!
			updatedAt: String!
		}
		type Location {
			latitude: Float!
			longitude: Float!
		}
		input ListProductInput{
			shopID: Int = 0
			page: Int = 1
			limit: Int = 10
		}
		input GetProductInput{
			id: ID = 0
			slugName: String = ""
		}
		input CreateProductInput {
			shopID: Int!
			name: String!
			quantity: Int!
			pricePerKG: Int!
			stockKG: Float!
			slugName: String!
			photo: Upload
		}
		input UpdateProductInput {
			id: Int!
			shopID: Int!
			name: String!
			quantity: Int!
			pricePerKG: Int!
			stockKG: Float!
			slugName: String!
			photo: Upload
		}
		input ListShopInput{
			page: Int = 1
			limit: Int = 10
		}
		input GetShopInput{
			id: ID = 0
			slugName: String = ""
		}
`
)
