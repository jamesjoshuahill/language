package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 5555, "port to listen on")
	flag.Parse()

	log.SetOutput(os.Stdout)
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
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("received '%s'\n", string(data))
}
