package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"livestream-exporter/utils"
)

func main() {
	// Health check
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println("livestream-exporter is health.")
		w.Header().Set("Content-Type", "application/json")
		response := map[string]bool{"up": true}
		json.NewEncoder(w).Encode(response)
	})

	// Register metrics collector
	liveCollector := utils.NewLiveCollector()
	prometheus.MustRegister(liveCollector)
	http.Handle("/metrics", promhttp.Handler())

	// Run server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
