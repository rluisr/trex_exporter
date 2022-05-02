package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rluisr/trex_exporter/trex"
	"log"
	"net/http"
)

var version string

func main() {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", config.ListenPort)
	log.Printf("Start exporter on %s/metrics", addr)

	http.HandleFunc("/", top)
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		probeHandler(w, r)
	})

	log.Fatal(http.ListenAndServe(addr, nil))
}

func probeHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	target := params.Get("target")
	if target == "" {
		http.Error(w, "Target parameter is missing", http.StatusBadRequest)
		return
	}

	worker := params.Get("worker")
	if worker == "" {
		http.Error(w, "Worker parameter is missing", http.StatusBadRequest)
	}

	fmt.Println("target: ", target)
	fmt.Println("worker: ", worker)

	registry := prometheus.NewRegistry()
	_ = trex.Probe(context.TODO(), target, worker, registry)

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

func top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version: %s", version)
}
