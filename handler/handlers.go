package handler

import (
	"encoding/json"
	"net/http"

	"app/manager"
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
