package main

import (
	"flag"
	"log"
	"os"

	"github.com/jamesjoshuahill/language/handler"
	"github.com/jamesjoshuahill/language/stats"
)

type Stats interface {
	Record(string)
	Summary() stats.Summary
}

func main() {
	var port, apiPort int
	flag.IntVar(&port, "port", 5555, "port to listen for natural language")
	flag.IntVar(&apiPort, "apiPort", 8080, "port to serve HTTP API")
	flag.Parse()

	log.SetOutput(os.Stdout)

	languageStats := stats.NewStats()

	listener := handler.New(languageStats)
	go listener.Listen(port)

	api := apiHandler{stats: languageStats}
	log.Fatal(api.ListenAndServe(apiPort))
}
