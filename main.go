package main

import (
	"context"
	"database/sql"
	_ "embed"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/dundunlabs/grapher"
	"github.com/graph-gophers/graphql-go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/vn-contrib/urhome-server/pkg/bar"
	"github.com/vn-contrib/urhome-server/resolver"
)

//go:embed schema.gql
var gql string

func main() {
	ctx := context.Background()

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL"))))
	db := bun.NewDB(sqldb, pgdialect.New()).WithQueryHook(bundebug.NewQueryHook(bundebug.FromEnv()))
	defer db.Close()

	hdr := http.NewServeMux()
	hdr.Handle("/graphql", &grapher.Handler{
		Schema: graphql.MustParseSchema(gql, resolver.New()),
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: hdr,
		BaseContext: func(l net.Listener) context.Context {
			return bar.ContextWithDB(ctx, db)
		},
	}
	go srv.ListenAndServe()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	srv.Shutdown(ctx)
}
