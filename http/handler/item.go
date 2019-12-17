package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gchumillas/crud-api/db/manager"
)

// GetItems handler.
func (env *Env) GetItems(w http.ResponseWriter, r *http.Request) {
	items := manager.GetItems(env.DB)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": items,
	})
}
