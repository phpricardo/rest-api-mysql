package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"encoding/json"
)

type Product struct {
	ID    string `json:"id"`
	Label string `json:"label"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "<user>:<senha>@tcp(127.0.0.1:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/products", getAllProducts).Methods("GET")

	http.ListenAndServe(":8000", router)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var products []Product
	result, err := db.Query("SELECT id, label from products")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var product Product
		err := result.Scan(&product.ID, &product.Label)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}
