package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Server struct {
	Port uint64
	Tags InstanceTags
}

func (s *Server) Start() {
	log.Printf("Listening on port %d\n", s.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", s.Port), s.NewMux()); err != nil {
		log.Fatalf("Error creating http server: %+v\n", err)
	}
	log.Printf("Exiting...\n")
}

func (s *Server) NewMux() *http.ServeMux {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", s.rootHandler)
	return mux
}

func (s *Server) rootHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// TODO: check the input, looking for DescribeTags

		io.WriteString(w, "hello, world!\n")
	}
}
