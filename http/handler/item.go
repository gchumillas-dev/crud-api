package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gchumillas/crud-api/db/manager"
)

// CreateItem handler.
func (env *Env) CreateItem(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		Title       string
		Description string
	}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&payload); err != nil {
		panic(err)
	}

	item := manager.Item{}
	item.Title = payload.Title
	item.Description = payload.Description
	item.CreateItem(env.DB)

	json.NewEncoder(w).Encode(map[string]interface{}{"id": item.ID})
}

// GetItems handler.
func (env *Env) GetItems(w http.ResponseWriter, r *http.Request) {
	items := manager.GetItems(env.DB)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": items,
	})
}
