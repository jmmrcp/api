package store

import (
	"encoding/json"
	"net/http"
)

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	products := c.Repository.GetProducts() // list of all products
	// log.Println(products)
	data, _ := json.Marshal(products)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
