package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jamesjoshuahill/language/stats"
)

type Stats interface {
	Record(string)
	Summary() stats.Summary
}

type apiHandler struct {
	port  int
	stats Stats
}

func (s apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(s.stats.Summary())
}

func (s apiHandler) Listen() error {
	mux := http.NewServeMux()
	mux.Handle("/stats", s)
	log.Printf("Serving HTTP on port %d...", s.port)
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), mux)
}
