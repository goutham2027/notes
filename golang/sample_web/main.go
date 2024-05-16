package main

import (
	"log"
	"net/http"
	"sample_web/handlers"
	sr "sample_web/subrouter"

	"github.com/gorilla/mux"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Register(r *mux.Router) {
	r.HandleFunc("/", handlers.HomeHandler)
	subrouter := r.PathPrefix("/home").Subrouter()
	subrouter.HandleFunc("/", sr.Index)
	r.HandleFunc("/get_id/{id}", handlers.ShowId)
}

func main() {
	r := mux.NewRouter()
	server := NewServer()
	server.Register(r)

	http.Handle("/", r)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())
}
