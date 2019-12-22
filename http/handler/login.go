package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gchumillas/crud-api/db/manager"
)

// Login handler.
func (env *Env) Login(w http.ResponseWriter, r *http.Request) {
	var body struct { Username string; Password string }
	parseBody(w, r, &body)

	// TODO: increase expiration time
	u := manager.NewUser()
	if !u.ReadUserByCredentials(env.DB, body.Username, body.Password) {
		httpError(w, forbiddenError)
		return
	}

	json.NewEncoder(w).Encode(u.NewToken(env.PrivateKey, env.Expiration))
}
