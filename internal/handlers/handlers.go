package handlers

import (
	"blockchain-node-gateway/api"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
)

/**
 * @notice handler functions below
 *
 */
func TxHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("content-type", "application/json")

	txRequest := new(api.SubmitTxRequest)
	err := json.NewDecoder(r.Body).Decode(&txRequest)

	/*
		if err != nil {
			http.Error(w, "invalid json to make tx", http.StatusBadRequest)
		}
		fmt.Printf("transaction received: %v", tx)
	*/

	switch err {
	case nil:
		http.Error(w, "invalid json to make tx", http.StatusBadRequest)
	default:
		fmt.Printf("transaction received: %v", txRequest.Tx)
	}

	resp := api.SubmitTxResponse{
		Status: "accepted",
	}
	json.NewEncoder(w).Encode(resp)
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK) // WriteHeader sends an HTTP response header with the provided status code.
	// w.Write([]uint8("node is alive"))
	/**
	resp := map[string]any{
		"status": "ok",
		"node":   "mini-node",
	}
	we are not using maps, instead we have request/response types defined
	*/
	resp := api.HealthResponse{
		Status: "ok",
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
