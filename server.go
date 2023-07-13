package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Server struct {
	listenAddr string
	db         *gorm.DB
}

func NewServer(listenAddr string, db *gorm.DB) *Server {
	return &Server{
		listenAddr: listenAddr,
		db:         db,
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
	var products []Product
	s.db.Find(&products)

	tmplData := map[string][]Product{
		"Products": products,
	}

	tmpl := template.Must(template.ParseFiles("templates/admin/list-products.html"))
	tmpl.Execute(w, tmplData)
}
