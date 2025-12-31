package main

import (
	handlers "blockchain-node-gateway/internal/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

var sec time.Duration = time.Second

func main() {
	l := log.New(os.Stdout, "blockchain-node", log.LstdFlags)

	// blockHandler := handlers.NewBlock(l)
	txHandler := handlers.NewTx(l)
	sm := mux.NewRouter()

	getTxRouter := sm.Methods(http.MethodGet).Subrouter()
	getTxRouter.HandleFunc("/tx", txHandler.ReceiveTx)

	s := http.Server{
		Addr:         ":8545",
		Handler:      sm,
		ReadTimeout:  5 * sec,
		WriteTimeout: 10 * sec,
		IdleTimeout:  12 * sec,
	}

	go func() {
		log.Println("starting server on localhost...")
		err := s.ListenAndServe()
		if err != nil {
			log.Printf("couldnt start server on localhost%s", s.Addr)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Printf("Get signal: %v", sig)

	ctx, _ := context.WithTimeout(context.Background(), 10*sec)
	err := s.Shutdown(ctx)
	if err != nil {
		log.Printf("couldnt shutdown gracefully: %s", err)
	}
}
