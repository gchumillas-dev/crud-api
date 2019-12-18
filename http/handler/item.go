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
	if err := item.CreateItem(env.DB); err != nil {
		panic(err)
	}

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
	if err := item.ReadItem(env.DB); err != nil {
		httpError(w, docNotFoundError)
		return
	}

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
	if err := item.UpdateItem(env.DB); err != nil {
		panic(err)
	}
}

// DeleteItem handler.
func (env *Env) DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemID, err := strconv.ParseInt(params["itemID"], 10, 32)
	if err != nil {
		panic(err)
	}

	item := manager.NewItem(itemID)
	if err := item.DeleteItem(env.DB); err != nil {
		panic(err)
	}
}

// GetItems handler.
func (env *Env) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := manager.GetItems(env.DB)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": items,
	})
}
