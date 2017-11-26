package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
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
