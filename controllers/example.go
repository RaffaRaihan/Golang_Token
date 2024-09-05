package controllers

import (
	"encoding/json"
	"net/http"
)

func (idb InDb) Example(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "Berhasil",
		"data":    nil,
	})
}

func (idb InDb) PingEndpoint(w http.ResponseWriter, r *http.Request) {
	err := idb.Ping()
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  false,
			"message": err.Error(),
			"data":    nil,
		})
	}

	json.NewEncoder(w).Encode("Pong")
}
