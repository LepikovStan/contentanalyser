package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"
)

type Upload struct {
	CHConn *sql.DB
}

type data struct {
	Title          string `json:"title"`
	Author         string `json:"author"`
	Language       string `json:"language"`
	URL            string `json:"url"`
	Referrer       string `json:"referrer"`
	TimeOnPage     int    `json:"time_on_page"`
	ScrollToMiddle bool   `json:"scroll_to_middle"`
	ScrollToEnd    bool   `json:"scroll_to_end"`
	Device         string `json:"device"`
	UserAgent      string `json:"user_agent"`
	Tags           string `json:"tags"`
	Keywords       string `json:"keywords"`
	Event          string `json:"event"`
}

func (h Upload) Handle(w http.ResponseWriter, r *http.Request) {
	d := data{}

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		panic([]byte(err.Error()))
	}

	args := []any{}
	args = append(
		args,
		1,
		time.Now(),
		time.Now(),
		d.Title,
		d.Author,
		d.Language,
		d.URL,
		d.Referrer,
		d.TimeOnPage,
		d.ScrollToMiddle,
		d.ScrollToEnd,
		d.Device,
		d.UserAgent,
		d.Event,
		1,
	)

	tx, err := h.CHConn.BeginTx(r.Context(), &sql.TxOptions{})
	if err != nil {
		panic([]byte(err.Error()))
	}

	tx.QueryRow("insert into rawdata_buffer_sum (userID, date, time, `title`, `author`, `language`, `url`, `referrer`, `time_on_page`, `scroll_to_middle`, `scroll_to_end`, `device`, `user_agent`, `event`, `count`) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", args...)

	if err := tx.Commit(); err != nil {
		panic([]byte(err.Error()))
	}

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Methods", "PUT, POST, GET, DELETE, PATCH, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "content-type")
	w.Write([]byte("ok"))
}
