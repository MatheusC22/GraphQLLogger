package main

import (
	"fmt"
	"goGRAPH/database"
	"goGRAPH/graph"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/streadway/amqp"
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

	conn, err := amqp.Dial("amqp://admin:secret@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	msgs, err := ch.Consume("MainQueue", "", true, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			parts := strings.Split(string(d.Body[:]), ";")
			database.UpdateLight(parts[0], parts[1])
		}
	}()
	fmt.Println("[*] - Waiting for Messages")
	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
	<-forever
}
