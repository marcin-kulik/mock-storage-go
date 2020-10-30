package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) store(w http.ResponseWriter, r *http.Request) {
	log.Print("Enter store")
	defer log.Print("Exit store")

	message := Message{}
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	s.Storage[message.Number] = ""
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	return
}

func (s *Server) getAll(w http.ResponseWriter, r *http.Request) {
	keys := make([]string, 0, len(s.Storage))
	for k := range s.Storage {
		keys = append(keys, k)
	}

	err := json.NewEncoder(w).Encode(keys)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

func alive(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("Service is alive")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	return
}
