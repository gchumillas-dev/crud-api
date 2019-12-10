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
