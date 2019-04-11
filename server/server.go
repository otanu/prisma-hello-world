package main

import (
	"log"
	"net/http"
	"os"
	prisma "prisma-hello-world/generated/prisma-client"
	"prisma-hello-world/gqlgen"

	"github.com/99designs/gqlgen/handler"
)

const defaultPort = "4000"

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = defaultPort
	}

	client := prisma.New(nil)
	resolver := gqlgen.Resolver{
		Prisma: client,
	}

	http.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	http.Handle("/query", handler.GraphQL(gqlgen.NewExecutableSchema(
		gqlgen.Config{Resolvers: &resolver})))

	log.Printf("Server is running on http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
