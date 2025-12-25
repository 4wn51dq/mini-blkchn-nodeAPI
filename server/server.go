package server

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK) // WriteHeader sends an HTTP response header with the provided status code.
	// w.Write([]uint8("node is alive"))
	resp := map[string]any{
		"status": "ok",
		"node":   "mini-node",
	}
	json.NewEncoder(w).Encode(resp)
}

func LatestBlockHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("content-type", "application/json")

	block := map[string]any{
		"hash":   "0x0676767696969r3t4rd",
		"txs":    rand.IntN(256),
		"height": 1,
	}

	json.NewEncoder(w).Encode(block)
}
