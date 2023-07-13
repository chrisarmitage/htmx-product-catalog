package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Run() {
	mux := chi.NewRouter()

	mux.Get("/admin/products", s.handleViewProducts)

	// start web server
	log.Println("Starting application on", s.listenAddr)

	err := http.ListenAndServe(fmt.Sprintf("%s", s.listenAddr), mux)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) handleViewProducts(w http.ResponseWriter, r *http.Request) {
	tmplData := map[string][]Product{
		"Products": {
			{Name: "Hand Weapon", Price: 7},
			{Name: "Dagger", Price: 4},
			{Name: "Spear", Price: 5},
		},
	}

	tmpl := template.Must(template.ParseFiles("templates/admin/list-products.html"))
	tmpl.Execute(w, tmplData)
}
