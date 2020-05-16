package cabaicatalog

const (
	Query = `
		# List Cabai Products
		searchProducts(params: ListProductInput!): [Product]
		productsByShop(params: ProductsByShopInput!): [Product]
		product(params: GetProductInput!): Product
		reviews(params: ListReviewInput!): [Review]
	`
	Mutation = `
		# Create new cabai product
		createProduct(params: CreateProductInput!): Product
		# update cabai product
		updateProduct(params: UpdateProductInput!): Product
		# delete cabai product
		deleteProduct(params: GetProductInput!): Success
		
		createReview(params: CreateReviewInput!): Review
		updateReview(params: UpdateReviewInput!): Review
		deleteReview(params: DeleteReviewInput!): Success
`
	Types = `
		scalar Upload
		type Review{
			id: ID!
			userID: ID!
			productID: ID!
			shopID: ID!
			title: String!
			content: String!
			photoURL: String!
			rating: Float!
			createdAt: String!
			updatedAt: String!
		}
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
			description: String!
			category: String!
			boughtKG: Float!
		}
		input ListReviewInput{
			productID: ID = "0"
			shopID: ID = "0"
		}
		input ListProductInput{
			category: String = ""
			search: String = ""
			province: String = ""
			orderBy: String = ""
			orderType: String = ""
			page: Int = 1
			limit: Int = 10
		}
		input ProductsByShopInput{	
			shopID: Int!
		}
		input GetProductInput{
			id: ID = 0
			slugName: String = ""
		}
		input CreateProductInput {
			shopID: Int!
			name: String!
			quantity: Int = 0
			pricePerKG: Int!
			stockKG: Float!
			slugName: String!
			photo: Upload
			description: String!
			category: String!
		}
		input UpdateProductInput {
			id: Int!
			shopID: Int!
			name: String!
			quantity: Int = 0
			pricePerKG: Int!
			stockKG: Float!
			slugName: String!
			description: String!
			photo: Upload
			category: String!
		}
		input CreateReviewInput{
			userID: String = "0"
			productID: String = "0"
			shopID: String = "0"
			title: String = ""
			content: String = ""
			photo: Upload
			rating: Float = 0.0
		}
		input UpdateReviewInput{
			id: String!
			userID: String = "0"
			productID: String = "0"
			shopID: String = "0"
			title: String = ""
			content: String = ""
			photo: Upload
			rating: Float = 0.0
		}
		input DeleteReviewInput{
			id: String!
			userID: String = "0"
		}
`
)
