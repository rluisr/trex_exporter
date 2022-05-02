package main

import (
	"fmt"
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

	collector := trex.NewCollector(config.TrexAPIAddress, config.TrexWorkerName)
	collector.StartCollect(config.CollectInterval)

	addr := fmt.Sprintf("0.0.0.0:%d", config.ListenPort)
	log.Printf("Start exporter on %s/metrics", addr)

	http.HandleFunc("/", top)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(addr, nil))
}

func top(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "version: %s", version)
}
