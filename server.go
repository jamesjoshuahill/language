package main

import (
	"flag"
	"log"
	"os"

	"github.com/jamesjoshuahill/language/stats"
)

func main() {
	var port, apiPort int
	flag.IntVar(&port, "port", 5555, "port to listen for natural language")
	flag.IntVar(&apiPort, "apiPort", 8080, "port to serve HTTP API")
	flag.Parse()

	log.SetOutput(os.Stdout)

	languageStats := stats.NewStats()

	listener := languageHandler{
		stats: languageStats,
		port:  port,
	}
	go listener.Listen()

	api := apiHandler{
		stats: languageStats,
		port:  apiPort,
	}
	log.Fatal(api.Listen())
}
