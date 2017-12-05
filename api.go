package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type apiHandler struct {
	stats Stats
}

func (s apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.stats.Summary())
}

func (s apiHandler) ListenAndServe(port int) error {
	mux := http.NewServeMux()
	mux.Handle("/stats", s)
	log.Printf("Serving HTTP on port %d...", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
}
