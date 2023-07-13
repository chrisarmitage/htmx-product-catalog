package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/oklog/ulid/v2"
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
	mux.Post("/admin/products/add", s.handleAddProduct)
	mux.Post("/admin/products/remove", s.handleRemoveProduct)

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

func (s *Server) handleAddProduct(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	// @TODO handle error
	price, _ := strconv.Atoi(r.PostFormValue("price"))

	product := Product{Id: ulid.Make(), Name: name, Price: price}
	s.db.Create(product)

	fmt.Println(name, price)

	tmpl := template.Must(template.ParseFiles("templates/admin/list-products.html"))
	tmpl.ExecuteTemplate(w, "product-list-element", product)
}

func (s *Server) handleRemoveProduct(w http.ResponseWriter, r *http.Request) {
	// @TODO handle error
	id, _ := ulid.Parse(r.PostFormValue("id"))
	fmt.Println(id.String())

	product := Product{Id: id}

	s.db.Delete(product)
}
