package auth

const (
	Query = `
		login(params: LoginInput!): Token
		refreshToken(params: RefreshTokenInput!): Token
		getUserInfo(token: String): User
	`
	Mutation = `
		register(params: RegisterInput!): Success
`
	Types = `
		type Success{
			success: Boolean!
		}
		type Token{
			token: String!
		}
		type User{
			id:      ID!
			name: String!
			userName: String!
			email: String!
			phone: String!
			type: Int!
			photoURL: String!
		}
		input LoginInput{
			userNameOrEmail: String!
			password: String!
		}
		input RefreshTokenInput{
			token: String!
		}
		input RegisterInput{
			userName: String!
			email: String!
			password: String!
			phoneNumber: String!
			fullName: String!
			role: Int!
			address: String
		}
`
)
