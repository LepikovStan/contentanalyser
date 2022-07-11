package handlers

import (
	"database/sql"
	"net/http"
)

type Healthcheck struct {
	CHConn *sql.DB
}

func (h Healthcheck) Handle(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, PATCH, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	w.Write([]byte("ok"))
}
