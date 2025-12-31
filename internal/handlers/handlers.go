package handlers

import (
	state "blockchain-node-gateway/internal/types"
	"log"
	"net/http"
)

type BlockHandler struct {
	l *log.Logger
}

type TxHandler struct {
	l *log.Logger
}

func NewBlock(l *log.Logger) *BlockHandler {
	return &BlockHandler{}
}

func NewTx(l *log.Logger) *TxHandler {
	return &TxHandler{}
}

func (th *TxHandler) ReceiveTx(w http.ResponseWriter, r *http.Request) {
	th.l.Println("Handle POST Tx")

	tx := &state.Tx{}
	err := tx.TxFromJSON(w, r)
	if err != nil {
		http.Error(w, "couldnt unmarshall json", http.StatusBadRequest)
		return
	}
	th.l.Printf("Received tx: %v", tx)
	state.RecordTx(tx)
}
