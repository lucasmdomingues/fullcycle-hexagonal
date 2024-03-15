package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/lucasmdomingues/hexagonal/adapters/db"
	"github.com/lucasmdomingues/hexagonal/application"
	"github.com/stretchr/testify/require"
)

func newConn() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	createTable(db)
	createProduct(db)

	return db
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE products (
		id string,
		name string,
		price float,
		status string
	);`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

func createProduct(db *sql.DB) {
	query := `INSERT INTO products(id,name,price,status) 
		VALUES('7403a96e-c4cf-4827-80fc-2141e5185a35','Product 1', 10.0, 'disabled');`

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

func TestProductDb_Get(t *testing.T) {
	conn := newConn()
	defer conn.Close()

	productDB := db.NewProductDB(conn)

	product, err := productDB.Get("7403a96e-c4cf-4827-80fc-2141e5185a35")
	require.NoError(t, err)
	require.Equal(t, "Product 1", product.GetName())
	require.Equal(t, 10.0, product.GetPrice())
	require.Equal(t, application.DISABLED, product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	conn := newConn()
	defer conn.Close()

	productDB := db.NewProductDB(conn)

	product := application.NewProduct()
	product.Name = "Product 1"
	product.Price = 25.0

	got, err := productDB.Save(product)
	require.NoError(t, err)
	require.Equal(t, product.GetName(), got.GetName())
	require.Equal(t, product.GetPrice(), got.GetPrice())
	require.Equal(t, product.GetStatus(), got.GetStatus())

	product.Status = application.ENABLED
	got, err = productDB.Save(product)
	require.NoError(t, err)
	require.Equal(t, product.GetStatus(), got.GetStatus())
}
