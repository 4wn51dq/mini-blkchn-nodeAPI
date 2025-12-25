package mempool

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Account struct {
	Address string
	Balance uint
}

type TX struct {
	ID     string
	from   string
	to     string
	amount uint
}

func TxHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("content-type", "application/json")

	tx := new(TX)
	err := json.NewDecoder(r.Body).Decode(&tx)
	if err != nil {
		http.Error(w, "invalid json to make tx", http.StatusBadRequest)
	}
	fmt.Printf("transaction received: %v", tx)

	w.WriteHeader(http.StatusAccepted)
}
