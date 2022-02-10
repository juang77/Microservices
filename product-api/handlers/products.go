package handlers

import (
	"Microservices/product-api/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		rx := regexp.MustCompile(`/([0-9]+)`)
		g := rx.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.getProducts(rw, r)
			return
		}

		if len(g[0]) != 2 {
			p.getProducts(rw, r)
			return
		}

		idString := g[0][1]

		id, err := strconv.Atoi(idString)
		if err == nil || id > 0 {
			p.getProductById(rw, r, id)
			return
		}

	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		//Expect the id in the URI
		rx := regexp.MustCompile(`/([0-9]+)`)
		g := rx.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]

		id, err := strconv.Atoi(idString)
		if err != nil || id < 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, rw, r)

	}

	//Handle an update

	//catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marchal json", http.StatusInternalServerError)
	}
}

func (p *Products) getProductById(rw http.ResponseWriter, r *http.Request, Id int) {
	lp := data.GetProductById(Id)
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marchal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Post Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Put Products")
	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}