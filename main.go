package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rluisr/trex_exporter/trex"
	"log"
	"net/http"
	"net/url"
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

	// target: http://hoge?worker=hoge
	// Need parse target for getting scheme://host and worker name
	target := params.Get("target")
	if target == "" {
		http.Error(w, "Target parameter is missing", http.StatusBadRequest)
		return
	}

	targetURL, err := url.Parse(target)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse target url: %s", targetURL), http.StatusBadRequest)
	}
	q := targetURL.Query()

	target = fmt.Sprintf("%s://%s", targetURL.Scheme, targetURL.Host)
	worker := q.Get("worker")
	if worker == "" {
		http.Error(w, "Worker parameter is missing", http.StatusBadRequest)
	}

	registry := prometheus.NewRegistry()
	_ = trex.Probe(context.TODO(), target, worker, registry)

	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
}

func top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version: %s", version)
}
