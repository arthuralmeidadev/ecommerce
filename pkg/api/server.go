package api

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	address string
	port    int
	handler http.Handler
}

func NewServer(address string, port int, handler http.Handler) *server {
	return &server{
		address,
		port,
		handler,
	}
}

func (s *server) Listen() {
	ctlChan := make(chan struct{})
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("%s:%d", s.address, s.port), s.handler)
		if err != nil {
			ctlChan <- struct{}{}
			log.Fatal("Couldn't start server: ", err)
		}
	}()
	log.Printf("\x1b[96mServer started successfully; listening on port %d \x1b[0m", s.port)
	<-ctlChan
}
