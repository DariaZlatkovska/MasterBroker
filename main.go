package main

import (
	"MasterBroker/config"
	"MasterBroker/storage"
	"MasterBroker/transport"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.InitRedis()
	config.InitMongo()

	transport.InitPrometheus()

	storage.InitMessageStorage()

	r := mux.NewRouter()
	r.Use(transport.MetricsMiddleware)
	r.HandleFunc("/messages", transport.PostMessageHandler).Methods("POST")
	r.HandleFunc("/messages", transport.GetMessageHandler).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())
	//r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusOK)
	//	w.Write([]byte("Broker is running!"))
	//})

	log.Println("Running broker on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
