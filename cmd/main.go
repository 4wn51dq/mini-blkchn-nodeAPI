package main

import (
	"blockchain-node-gateway/mempool"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
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

func latestBlockHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	// r:= chi.NewRouter()
	// server.Handler(r)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/block", latestBlockHandler)
	http.HandleFunc("/tx", mempool.TxHandler)
	fmt.Println("Starting mini blockchain node...")
	http.ListenAndServe("localhost:8545", nil) // nil is Go's defualt router
}
