package db_test

import (
	"database/sql"
	"github.com/almeidacavalcante/ports-and-adapters/adapters/db"
	"github.com/almeidacavalcante/ports-and-adapters/application"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
	"testing"
)

func setup(t *testing.T) *sql.DB {
	database, err := sql.Open("sqlite3", ":memory:")
	require.NoError(t, err)

	err = createTable(database)
	require.NoError(t, err)

	err = createProduct(database)
	require.NoError(t, err)

	return database
}

func createTable(db *sql.DB) error {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"status" string,
		"price" float
	);`
	_, err := db.Exec(table)
	return err
}

func createProduct(db *sql.DB) error {
	insert := `INSERT INTO products values("abc", "Product Test", "disabled", 0.0)`
	_, err := db.Exec(insert)
	return err
}

func TestProductDb_Get(t *testing.T) {
	database := setup(t)
	defer database.Close()

	productDb := db.NewProductDb(database)
	product, err := productDb.Get("abc")
	require.Nil(t, err)
	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	database := setup(t)
	defer database.Close()
	productDb := db.NewProductDb(database)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Status = application.DISABLED
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "Product Test", productResult.GetName())
	require.Equal(t, 25.0, productResult.GetPrice())
	require.Equal(t, "disabled", productResult.GetStatus())

	product.Status = application.ENABLED
	product.Price = 50
	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, "Product Test", productResult.GetName())
	require.Equal(t, 50.0, productResult.GetPrice())
	require.Equal(t, "enabled", productResult.GetStatus())

}
