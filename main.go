package main

import (
    "log"
    "encoding/json"
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"

    "live-supplier-exporter/config"
    "live-supplier-exporter/utils"
)


func main() {
    // Loads config file
    if err := config.LoadConfig(); err != nil {
	log.Fatalf("Error loading config: %s", err)
    }

    // Goroutine for watches config file
    go config.WatchConfig()

    // Get config
    log.Printf("config contenxt: %s\n", config.AppConfig)

    // Health check
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
        //log.Println("live-supplier-exporter is health.")

        w.Header().Set("Content-Type", "application/json")
        response := map[string]bool{"up": true}
        json.NewEncoder(w).Encode(response)
    })

    // Register metrics collector
    liveCollector := utils.NewLiveCollector()
    prometheus.MustRegister(liveCollector)
    // Metrics api
    http.Handle("/metrics", promhttp.Handler())

    // Run server
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
