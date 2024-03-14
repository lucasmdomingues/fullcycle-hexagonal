package main

import (
	"database/sql"
	"log"

	productDB "github.com/lucasmdomingues/hexagonal/adapters/db"
	"github.com/lucasmdomingues/hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	productDBAdapter := productDB.NewProductDB(db)
	productService := application.NewProductService(productDBAdapter)

	product, err := productService.Create("Product 1", 10)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", product)

	product, err = productService.Enable(product)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", product)
}
