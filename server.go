package main

import (
	"goGRAPH/graph"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

const defaultPort = "3333"

func main() {
	//queue := queue.NewRabbitMQService()
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	http.ListenAndServe(":"+port, nil)

	// msgs, err := queue.Ch.Consume("MainQueue", "", true, false, false, false, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// forever := make(chan bool)

	// go func() {
	// 	for d := range msgs {
	// 		parts := strings.Split(string(d.Body[:]), ";")
	// 		database.UpdateEndpoint(parts[0], parts[1])
	// 	}
	// }()
	// fmt.Println("[*] - Waiting for Messages")
	// <-forever
}
