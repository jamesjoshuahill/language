package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/jamesjoshuahill/language/stats"
)

type languageHandler struct {
	port  int
	stats Stats
}

func (l languageHandler) Listen() {
	log.Printf("Listening on port %d...", l.port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", l.port))
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go l.connHandler(conn)
	}
}

func (l languageHandler) connHandler(conn net.Conn) {
	defer conn.Close()
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	language := string(data)
	log.Printf("received '%s'\n", language)
	l.stats.Record(language)
}

type Stats interface {
	Record(string)
	Summary() stats.Summary
}

type statsHandler struct {
	stats Stats
}

func (h statsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(h.stats.Summary())
}

func main() {
	var port, webPort int
	flag.IntVar(&port, "port", 5555, "port to listen for natural language")
	flag.IntVar(&webPort, "webPort", 8080, "port to serve HTTP endpoints")
	flag.Parse()

	log.SetOutput(os.Stdout)

	languageStats := stats.NewStats()

	listener := languageHandler{
		stats: languageStats,
		port:  port,
	}
	go listener.Listen()

	mux := http.NewServeMux()
	mux.Handle("/stats", statsHandler{languageStats})
	log.Printf("Serving HTTP on port %d...", webPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", webPort), mux))
}
