package main

import (
	handlers "blockchain-node-gateway/internal/handlers"
	"fmt"
	"net/http"
)

func main() {
	// r:= chi.NewRouter()
	// server.Handler(r)
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/block", handlers.LatestBlockHandler)
	http.HandleFunc("/tx", handlers.TxHandler)
	fmt.Println("Starting mini blockchain node...")
	http.ListenAndServe("127.0.0.1:3001", nil) // nil is Go's defualt router
}
