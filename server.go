package main

import (
	"github.com/PrinceNarteh/gqlgen_cc/pkg/repository"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/PrinceNarteh/gqlgen_cc/graph"
	"github.com/PrinceNarteh/gqlgen_cc/graph/generated"
)

const defaultPort = "8080"

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("could not load .env file")
	}

	// Initializing database
	repository.InitDAO()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
