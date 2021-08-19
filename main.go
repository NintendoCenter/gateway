package main

import (
	"log"
	"net/http"

	"NintendoCenter/gateway/graph/generated"
	"NintendoCenter/gateway/graph/resolvers"
	"NintendoCenter/gateway/internal/protos"
	"google.golang.org/grpc"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)


func main() {
	cfg, _ := NewConfig()

	conn, err := grpc.Dial(cfg.CollectionAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("error on creating grpc connection", err)
	}

	client := protos.NewGameCollectionClient(conn)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolvers.NewResolver(client)}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/"))
	http.Handle("/", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.WebPort)
	log.Fatal(http.ListenAndServe(":"+cfg.WebPort, nil))
}
