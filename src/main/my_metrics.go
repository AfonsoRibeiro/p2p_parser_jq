package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "p2p_parser_jq_TODO",
		Help: "The total number of processed events",
	})
)

func setup_prometheus(prometheusport uint) {
	http.Handle("/", promhttp.Handler())

	go http.ListenAndServe(fmt.Sprintf(":%d", prometheusport), nil)
}
