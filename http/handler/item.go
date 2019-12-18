package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gchumillas/crud-api/db/manager"
	"github.com/gorilla/mux"
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

// ReadItem handler.
func (env *Env) ReadItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemID, err := strconv.ParseInt(params["itemID"], 10, 32)
	if err != nil {
		panic(err)
	}

	item := manager.NewItem(itemID)
	item.ReadItem(env.DB)

	json.NewEncoder(w).Encode(item)
}

// UpdateItem handler.
func (env *Env) UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemID, err := strconv.ParseInt(params["itemID"], 10, 32)
	if err != nil {
		panic(err)
	}

	var payload struct {
		Title       string
		Description string
	}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&payload); err != nil {
		panic(err)
	}

	item := manager.NewItem(itemID)
	item.Title = payload.Title
	item.Description = payload.Description
	item.UpdateItem(env.DB)
}

// GetItems handler.
func (env *Env) GetItems(w http.ResponseWriter, r *http.Request) {
	items := manager.GetItems(env.DB)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": items,
	})
}
