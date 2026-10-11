package main

import (
	"context"
	_ "embed"
	"net/http"
	"os"
	"os/signal"

	"github.com/dundunlabs/grapher"
	"github.com/graph-gophers/graphql-go"
	"github.com/vn-contrib/urhome-server/resolver"
)

//go:embed schema.gql
var gql string

func main() {
	ctx := context.Background()

	hdr := http.NewServeMux()
	hdr.Handle("/graphql", &grapher.Handler{
		Schema: graphql.MustParseSchema(gql, resolver.New()),
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: hdr,
	}
	go srv.ListenAndServe()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	srv.Shutdown(ctx)
}
