package main

import (
    "log"
    "net/http"

    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)


func main() {

    // Health check
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
        log.Printf("Return: %v", w)
    })

    // register metrics collector
    c := NewDataCollector()
    prometheus.MustRegister(c)

    // metrics api
    http.Handle("/metrics", promhttp.Handler())

    // Run server
    log.Println("Starting server on :9097")
    if err := http.ListenAndServe(":9097", nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}
