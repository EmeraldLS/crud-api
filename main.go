package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	ID            string        `json:"id"`
	Name          string        `json:"name"`
	Price         float64       `json:"price"`
	Seller        *Seller       `json:"seller,omitempty"`
	ProductDetail ProductDetail `json:"product_details"`
}

type ProductDetail struct {
	SerialNumber string   `json:"serial_number"`
	Alternatives []string `json:"product_alternatives,omitempty"`
}

type Seller struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

var products []Product

func main() {

	products = append(products, Product{
		ID: "1", Name: "Rice", Price: 44.3, Seller: &Seller{
			FirstName: "Lawrence",
			LastName:  "Oluwasegun",
		},
		ProductDetail: ProductDetail{
			SerialNumber: strconv.Itoa(rand.Intn(1000000000000)),
			Alternatives: []string{"Garri", "Salad", "Pure Water", "Buns"},
		},
	})
	products = append(products, Product{
		ID: "2", Name: "Beans", Price: 12.2, Seller: &Seller{
			FirstName: "Sanni",
			LastName:  "Abdullah",
		},
		ProductDetail: ProductDetail{
			SerialNumber: strconv.Itoa(rand.Intn(1000000000000)),
			Alternatives: []string{"Egg", "Egusi", "Vegetable", "Milk"},
		},
	})
	crudAPI(":8000")
}

func crudAPI(port string) {
	r := mux.NewRouter()
	r.HandleFunc("/products", getAllProduct).Methods("GET")
	r.HandleFunc("/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/products", addProduct).Methods("POST")
	r.HandleFunc("/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/products/{id}", deleteProduct).Methods("DELETE")

	fmt.Printf("Development server started at port %v", port)
	log.Fatal(http.ListenAndServe(port, r))

}

func getAllProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range products {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func addProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	product.ID = strconv.Itoa(rand.Intn(100000))
	product.ProductDetail.SerialNumber = strconv.Itoa(rand.Intn(1000000000000))
	products = append(products, product)
	json.NewEncoder(w).Encode(products)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			var product Product
			_ = json.NewDecoder(r.Body).Decode(&product)
			product.ID = strconv.Itoa(rand.Intn(1000000))
			product.ProductDetail.SerialNumber = strconv.Itoa(rand.Intn(1000000000000))
			products = append(products, product)
			json.NewEncoder(w).Encode(products)
			break
		}
	}

}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range products {
		if item.ID == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
	}
}
