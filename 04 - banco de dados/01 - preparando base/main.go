package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// temos que importar o driver mysql
	product := NewProduct("produto Teste", 9999.90)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Price = 8888.88
	fmt.Println(product)

	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selectProductById(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\n")
	fmt.Printf("Product select: %v, tem valor de: %.2f\n ", p.Name, p.Price)
	fmt.Printf("\n")

	productListBefore, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}

	for _, p := range productListBefore {
		fmt.Printf("Product before: %v, tem valor de: %.2f\n", p.Name, p.Price)
	}

	err = deleteProductById(db, product.ID)
	if err != nil {
		panic(err)
	}

	productListAfter, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")
	fmt.Printf("\n")

	for _, p := range productListAfter {
		fmt.Printf("Product After: %v, tem valor de: %.2f\n", p.Name, p.Price)
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}

	return nil
}

func selectProductById(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare(`
		select id, name, price from products
		where id = ?
	`)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	var p Product

	//buscando apenas uma linha
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	rows, err := db.Query(`select id, name, price from products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product

	for rows.Next() {
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func deleteProductById(db *sql.DB, id string) error {
	stmt, err := db.Prepare(`
		delete from products
		where id = ?
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
	// buscar algo -> query
	// executar algo -> exec
}
