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

func main() {
	var port, webPort int
	flag.IntVar(&port, "port", 5555, "port to listen for natural language")
	flag.IntVar(&webPort, "webPort", 8080, "port to serve HTTP endpoints")
	flag.Parse()

	log.SetOutput(os.Stdout)

	go languageListener(port)

	mux := http.NewServeMux()
	mux.HandleFunc("/stats", statsHandler)
	log.Printf("Serving HTTP on port %d...", webPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", webPort), mux))
}

func languageListener(port int) {
	log.Printf("Listening on port %d...", port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go connHandler(conn)
	}
}

func connHandler(conn net.Conn) {
	defer conn.Close()
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("received '%s'\n", string(data))
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(stats.Summary{
		Count:       5,
		Top5Words:   []string{"here", "are", "some", "more", "words"},
		Top5Letters: []string{"e", "r", "o", "s", "h"},
	})
}
