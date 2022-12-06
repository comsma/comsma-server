package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/comsma/comsma/graph"
	"github.com/comsma/comsma/graph/generated"
	"github.com/comsma/comsma/internal/gallery"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8082"

func main() {

	gallery.Init()
	fmt.Println(gallery.GetGalleries())

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
