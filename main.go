package main

import (
	"fmt"
	"github.com/idzharbae/cabai-gqlserver/gql/auth"
	authfetcher "github.com/idzharbae/cabai-gqlserver/gql/auth/fetcher/grpc"
	authmutator "github.com/idzharbae/cabai-gqlserver/gql/auth/mutator/grpc"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog"
	"github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/fetcher/grpc"
	grpcmutator "github.com/idzharbae/cabai-gqlserver/gql/cabaicatalog/mutator/grpc"
	"github.com/idzharbae/cabai-gqlserver/middleware"
	"github.com/idzharbae/marketplace-backend/svc/auth/authproto"
	"github.com/idzharbae/marketplace-backend/svc/catalog/catalogproto"
	"github.com/idzharbae/marketplace-backend/svc/resources/protoresources"
	"google.golang.org/grpc"
	"log"
	"net/http"

	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
)

var schema *graphql.Schema

type Handler struct {
	*cabaicatalog.CabaiCatalogHandler
	*auth.AuthHandler
}

func NewHandler() *Handler {
	opts := []grpc.DialOption{grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(12582912))}
	catalogConn, err := grpc.Dial("127.0.0.1:1445", opts...)
	if err != nil {
		panic(err)
	}
	resourcesConn, err := grpc.Dial("127.0.0.1:1442", opts...)
	if err != nil {
		panic(err)
	}
	catalogHandler := getCatalogHandler(catalogConn, resourcesConn)

	authConn, err := grpc.Dial("127.0.0.1:1444", opts...)
	authHandler := getAuthHandler(authConn, catalogConn)

	return &Handler{CabaiCatalogHandler: catalogHandler, AuthHandler: authHandler}
}

func getAuthHandler(conn *grpc.ClientConn, catalogConn *grpc.ClientConn) *auth.AuthHandler {
	authConn := authproto.NewMarketplaceAuthClient(conn)
	catalog := catalogproto.NewMarketplaceCatalogClient(catalogConn)

	tokenFetcher := authfetcher.NewTokenFetcher(authConn)
	userMutator := authmutator.NewUserMutator(authConn, catalog)
	userReader := authfetcher.NewUserReader(authConn)
	return auth.NewAuthHandler(tokenFetcher, userMutator, userReader)
}

func getCatalogHandler(catalog, resources *grpc.ClientConn) *cabaicatalog.CabaiCatalogHandler {
	catalogConn := catalogproto.NewMarketplaceCatalogClient(catalog)
	resourcesConn := protoresources.NewMarketplaceResourcesClient(resources)

	productReader := grpcfetcher.NewProductReader(catalogConn)
	productWriter := grpcmutator.NewProductWriter(catalogConn, resourcesConn)
	shopReader := grpcfetcher.NewShopReader(catalogConn)
	catalogHandler := cabaicatalog.NewCabaiCatalogHandler(productReader, productWriter, shopReader)
	return catalogHandler
}

func NewSchemaSring() string {
	schemaString := fmt.Sprintf(`
		schema {
			query: Query
			mutation: Mutation
		}
		# List Cabai Products
		type Query{
			%s
			%s
		}
		type Mutation{
			%s
			%s
		}
	`, cabaicatalog.Query, auth.Query,
		cabaicatalog.Mutation, auth.Mutation)
	types := cabaicatalog.Types + auth.Types
	return schemaString + types
}

func init() {
	handler := NewHandler()
	schema = graphql.MustParseSchema(NewSchemaSring(), handler)
}

func main() {
	http.Handle("/", middleware.CorsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	})))

	http.Handle("/query", middleware.CorsMiddleware(&relay.Handler{Schema: schema}))
	fmt.Println("Listening to port 4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.min.css" rel="stylesheet" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/es6-promise/4.1.1/es6-promise.auto.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/2.0.3/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/16.2.0/umd/react.production.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react-dom/16.2.0/umd/react-dom.production.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.11.11/graphiql.min.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				return fetch("/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}

			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
