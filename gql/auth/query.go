package auth

const (
	Query = `
		login(params: LoginInput!): Token
		refreshToken(params: RefreshTokenInput!): Token
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
		}
`
)
