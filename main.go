package main

import (
	"context"
	"database/sql"
	_ "embed"
	"net"
	"net/http"
	"os"
	"os/signal"

	firebase "firebase.google.com/go"
	"github.com/dundunlabs/grapher"
	"github.com/graph-gophers/graphql-go"
	_ "github.com/joho/godotenv/autoload"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
	"github.com/vn-contrib/urhome-server/contexts"
	"github.com/vn-contrib/urhome-server/handler"
	"github.com/vn-contrib/urhome-server/pkg/bar"
	"github.com/vn-contrib/urhome-server/resolver"
	"google.golang.org/api/option"
)

//go:embed schema.gql
var gql string

func main() {
	ctx := context.Background()

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(os.Getenv("DATABASE_URL"))))
	db := bun.NewDB(sqldb, pgdialect.New()).WithQueryHook(bundebug.NewQueryHook(bundebug.FromEnv()))
	defer db.Close()

	fbApp, err := firebase.NewApp(ctx, nil, option.WithAuthCredentialsFile(option.ServiceAccount, os.Getenv("FIREBASE_CREDENTIALS_FILE")))
	if err != nil {
		panic(err)
	}
	fbAuth, err := fbApp.Auth(ctx)
	if err != nil {
		panic(err)
	}

	hdr := http.NewServeMux()
	hdr.Handle("/graphql", &handler.AuthHandler{
		Next: &grapher.Handler{
			Schema: graphql.MustParseSchema(gql, resolver.New(), graphql.UseFieldResolvers()),
		},
	})

	srv := http.Server{
		Addr:    ":8080",
		Handler: hdr,
		BaseContext: func(l net.Listener) context.Context {
			httpCtx := bar.ContextWithDB(ctx, db)
			httpCtx = contexts.WithFirebaseAuth(httpCtx, fbAuth)
			return httpCtx
		},
	}
	go srv.ListenAndServe()

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	srv.Shutdown(ctx)
}
