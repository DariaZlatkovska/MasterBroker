package main

import (
	"MasterBroker/config"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Inicjalizacja Redis i MongoDB
	config.InitRedis()
	config.InitMongo()

	// Router HTTP
	r := mux.NewRouter()
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Broker działa!"))
	})

	// Start serwera
	log.Println("Uruchamianie brokera na porcie 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
