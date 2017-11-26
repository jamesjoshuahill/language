package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 5555
	fmt.Printf("Starting server on port %d...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
