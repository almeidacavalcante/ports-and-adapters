package main

import (
	"database/sql"
	"fmt"
	"github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	database, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return
	}

	if err = database.Ping(); err != nil {
		log.Fatalf("Database is not accessible: %v", err)
		return
	}

	productDbAdapter := db.NewProductDb(database)
	if productDbAdapter == nil {
		fmt.Println("Failed to create ProductDbAdapter")
		return
	}

	productService := application.NewProductService(productDbAdapter)
	if productService == nil {
		fmt.Println("Failed to create ProductService")
		return
	}

	product, err := productService.Create("Banana", 9)
	productService.Enable(product)

	id := product.GetID()
	product, err = productService.Get(id)

}
