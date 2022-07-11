package main

import (
	"database/sql"
	"net/http"
	"time"

	clickhouse "github.com/ClickHouse/clickhouse-go/v2"
	"github.com/LepikovStan/contentanalyser/httpsrv"
	"github.com/LepikovStan/contentanalyser/httpsrv/handlers"
	"github.com/gorilla/mux"
)

func main() {
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "",
		},
		// TLS: &tls.Config{
		// 	InsecureSkipVerify: true,
		// },
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		Compression: &clickhouse.Compression{
			clickhouse.CompressionLZ4,
		},
		Debug: true,
	})
	conn.SetMaxIdleConns(5)
	conn.SetMaxOpenConns(10)
	conn.SetConnMaxLifetime(time.Hour)

	if err := conn.Ping(); err != nil {
		panic(err)
	}

	r := routes(conn)
	httpsrv.Start("0.0.0.0", 80, r)
}

func routes(chConn *sql.DB) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", H{}.Handle)

	r.HandleFunc(
		"/hc",
		handlers.Healthcheck{
			CHConn: chConn,
		}.Handle,
	).Methods(http.MethodGet)

	r.HandleFunc(
		"/upload",
		handlers.Upload{
			CHConn: chConn,
		}.Handle,
	).Methods(http.MethodPost)

	r.HandleFunc(
		"/upload",
		handlers.Access{
			CHConn: chConn,
		}.Handle,
	).Methods(http.MethodOptions)

	return r
}

type H struct{}

func (h H) Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
