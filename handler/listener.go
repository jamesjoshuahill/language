package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
)

//go:generate counterfeiter -o fakes/recorder.go . Recorder
type Recorder interface {
	Record(string)
}

//go:generate counterfeiter -o fakes/conn.go . Conn
type Conn interface {
	io.Reader
	io.Closer
}

type languageHandler struct {
	stats Recorder
}

func New(stats Recorder) *languageHandler {
	return &languageHandler{stats: stats}
}

func (l languageHandler) Listen(port int) {
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
		go l.Handle(conn)
	}
}

func (l languageHandler) Handle(conn Conn) {
	defer conn.Close()
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal(err)
	}

	language := string(data)
	log.Printf("received '%s'\n", language)
	l.stats.Record(language)
}
