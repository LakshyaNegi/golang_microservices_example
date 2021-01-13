package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/LakshyaNegi/golang_microservices_example/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {
// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(g) != 1 {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 		}

// 		if len(g[0]) != 2 {
// 			http.Error(rw, "Invalid URI", http.StatusBadRequest)
// 		}

// 		idStr := g[0][1]
// 		id, _ := strconv.Atoi(idStr)

// 		p.l.Printf("ID : %v\n", id)

// 		p.updateProduct(id, rw, r)
// 		return

// 	}
// 	//catch all
// 	rw.WriteHeader(http.StatusMethodNotAllowed)
// }

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal products", http.StatusInternalServerError)
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	// err := prod.FromJSON(r.Body)
	// if err != nil {
	// 	http.Error(rw, "Unable to unmarshal products", http.StatusInternalServerError)
	// }

	p.l.Printf("PROD %#v", prod)

	data.AddProduct(&prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idstr, _ := vars["id"]
	id, _ := strconv.Atoi(idstr)
	// if err != nil {
	// 	http.Error(rw, "Unable to find id", http.StatusInternalServerError)
	// }

	p.l.Println("Handle POST")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product Not Found", http.StatusBadRequest)
	}
	if err != nil {
		http.Error(rw, "Product Not Found", http.StatusInternalServerError)
	}
}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}

		err := prod.FromJSON(r.Body)
		if err != nil {
			http.Error(rw, fmt.Sprint("e", err), http.StatusInternalServerError)
			return
		}

		err = prod.Validate()
		if err != nil {
			http.Error(rw, "Unable to validate product", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, *prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}
