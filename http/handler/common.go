package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Env contains common variables, such as the database access, etc.
type Env struct {
	DB          *sql.DB
	PrivateKey  string
	Expiration  time.Duration
	RowsPerPage int
}

// Common HTTP status errors.
type httpStatus struct {
	code int
	msg  string
}

var (
	badRequestError   = httpStatus{400, "Bad request"}
	docNotFoundError  = httpStatus{404, "Document not found"}
	unauthorizedError = httpStatus{401, "Not authorized"}
	forbiddenError    = httpStatus{403, "Forbidden"}
)

func httpError(w http.ResponseWriter, status httpStatus) {
	http.Error(w, status.msg, status.code)
	log.Printf("http error (%d): %s", status.code, status.msg)
	return
}

func parseBody(w http.ResponseWriter, r *http.Request, body interface{}) {
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(body); err != nil {
		panic(err)
	}
}

func getParam(r *http.Request, key, def string) (param string) {
	if param = r.FormValue(key); len(param) == 0 {
		param = def
	}

	return
}

func inArray(item string, items []string) bool {
	for index := range items {
		if items[index] == item {
			return true
		}
	}

	return false
}
