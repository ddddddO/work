package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"

	myhandler "github.com/ddddddO/work/graphql/handler"
)

func main() {
	// neo4j
	driver, err := newNeo4j()
	if err != nil {
		log.Fatal(err)
	}
	defer driver.Close()

	// graphql
	gh := myhandler.NewGraphqlHandler(driver)
	graphqlHandler, err := gh.Handler()
	if err != nil {
		log.Fatal(err)
	}

	// health
	hh := myhandler.NewHealthHandler(driver)
	// debug
	dh := myhandler.NewDebugHandler(driver)

	mux := http.NewServeMux()
	mux.Handle("/graphql", graphqlHandler)
	mux.HandleFunc("/health", hh.Health)
	mux.HandleFunc("/debugNeo4j", dh.Debug)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("start")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)
	<-sig

	// graceful shutdown...
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

func newNeo4j() (neo4j.Driver, error) {
	const dsn = "neo4j://localhost:7687"
	driver, err := neo4j.NewDriver(dsn, neo4j.BasicAuth("username", "password", ""))
	if err != nil {
		return nil, err
	}
	return driver, nil
}
