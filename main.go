package main

import (
	"log"
	"net/http"
	"os"

	"NintendoCenter/gateway/graph/generated"
	"NintendoCenter/gateway/graph/resolvers"
	"NintendoCenter/gateway/internal/protos"
	"google.golang.org/grpc"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		log.Fatal("error on creating grpc connection", err)
	}

	client := protos.NewGameCollectionClient(conn)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers.NewResolver(client)}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	http.Handle("/", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
