package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// categorySomEVideo := Category{Name: "Som e Vídeo"}
	// db.Create(&categorySomEVideo)

	// categoryInformatica := Category{Name: "Informática"}
	// db.Create(&categoryInformatica)

	// productCaixaDeSom := Product{
	// 	Name:       "Caixinha de som",
	// 	Price:      350.6,
	// 	Categories: []Category{categoryInformatica, categorySomEVideo},
	// }
	// db.Create(&productCaixaDeSom)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		fmt.Println(category.Name, ":")
		for _, product := range category.Products {
			fmt.Println("- ", product.Name)
		}
	}
}
