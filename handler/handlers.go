package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"app/manager"
	"app/str"
)

// Login handler.
func (env *Env) Login(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Username string
		Password string
	}
	parseBody(w, r, &body)

	u := manager.NewUser()
	if !u.ReadUserByCredentials(env.DB, body.Username, body.Password) {
		httpError(w, forbiddenError)
		return
	}

	json.NewEncoder(w).Encode(u.NewToken(env.PrivateKey, env.Expiration))
}

// Home handler.
func (env *Env) Home(w http.ResponseWriter, r *http.Request) {
	u := getUser(r)
	username := str.Ucfirst(u.Username)

	msg := fmt.Sprintf(
		"Hi %s, welcome to CMSystem. Select a section from the main menu to start editing.",
		username)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title":   "Welcome to CMSystem",
		"message": msg,
	})
}

// SectionNotes handler.
func (env *Env) SectionNotes(w http.ResponseWriter, r *http.Request) {
	msg := "Lorem ipsum dolor sit amet..."
	json.NewEncoder(w).Encode(map[string]interface{}{
		"title":   nil,
		"message": msg,
	})
}

// Sections handler.
func (env *Env) Sections(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]section{
		section{"1", "Section 1", []section{
			section{"11", "Section 11", []section{
				section{"111", "Section 111", []section{}},
			}},
			section{"12", "Section 12", []section{}},
		}},
		section{"2", "Section 2", []section{
			section{"21", "Section 21", []section{}},
		}},
		section{"3", "Section 3", []section{}},
	})
}
