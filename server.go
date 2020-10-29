package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Server struct {
	Router  *mux.Router
	Storage map[string]string
}

func NewServer() *Server {
	return &Server{
		Router:  mux.NewRouter(),
		Storage: make(map[string]string),
	}
}

func (s *Server) setHandlers() {
	s.Router.HandleFunc("/store", s.store).Methods("POST")
	s.Router.HandleFunc("/getAll", s.getAll).Methods("GET")
	s.Router.HandleFunc("/alive", alive).Methods("GET")
}
func Run() {
	s := NewServer()
	s.setHandlers()
	log.Printf("Starting server on %s", ":5010")
	err := http.ListenAndServe(":5010", s.Router)
	log.Fatal(err)
}
