package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	// * create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 2000.90,
	// })

	// * create batch
	// products := []Product{
	// 	{Name: "Notebook 2", Price: 3000.00},
	// 	{Name: "Mouse", Price: 300.00},
	// 	{Name: "Gabinete", Price: 1000.00},
	// }
	// db.Create(&products)

	// * select one
	var product Product
	var productMouse Product

	db.First(&product, 2)
	fmt.Println(product)
	fmt.Println("\n")

	db.First(&productMouse, "name = ?", "Mouse")
	fmt.Println(productMouse)
	fmt.Println("\n")

	// * select all
	// var products []Product
	// db.Find(&products)

	// for _, item := range products {
	// 	fmt.Println(item)
	// }
	// fmt.Println("\n")

	// * limit
	// var products []Product
	// db.Limit(2).Find(&products)
	// for _, item := range products {
	// 	fmt.Println(item)
	// }

	// * Offset
	// fmt.Println("\n")
	// db.Limit(2).Offset(4).Find(&products)
	// for _, item := range products {
	// 	fmt.Println(item)
	// }

	// * where
	var products []Product
	// db.Where("price >= ?", 1000).Find(&products)
	// for _, item := range products {
	// 	fmt.Println(item)
	// }

	db.Where("name like ?", "%book%").Where("price > ?", 2500).Find(&products)
	for _, item := range products {
		fmt.Println(item)
	}
}
