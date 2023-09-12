package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           int `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   int
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

// * gorm.Model
// permite que gorm gerencie algumas informações da tabela como createdAt, updatedAt...
// habilita softDelete

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product
}

// ? []Product -> relacionamento "um pra muitos"

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// * create
	// db.Create(&Product{
	// 	Name:  "Notebook",
	// 	Price: 2000.90,
	// })

	// eletronicosCategory := Category{Name: "Eletronicos"}
	// somVideoCategory := Category{Name: "Som E Vídeo"}
	// acessoriosCategory := Category{Name: "Acessórios"}
	// db.Create(&somVideoCategory)
	// db.Create(&acessoriosCategory)
	// db.Create(&eletronicosCategory)

	// * create batch
	// products := []Product{
	// 	{Name: "Notebook 2", Price: 3000.00},
	// 	{Name: "Mouse", Price: 300.00},
	// 	{Name: "Gabinete", Price: 1000.00},
	// }
	// db.Create(&products)

	// categories := []Category{
	// 	{Name: "Eletronicos"},
	// 	{Name: "Acessórios internos"},
	// 	{Name: "Acessórios Externos"},
	// }
	// db.Create(&categories)

	// * create com relacionamento "um pra um"
	// hasOne ou "um pra um"
	// productCaixaDeSom := Product{
	// 	Name:       "Caixa de som",
	// 	Price:      400.00,
	// 	CategoryID: 2,
	// }
	// db.Create(&productCaixaDeSom)

	// productNotebook := Product{
	// 	Name:       "Notebook",
	// 	Price:      2000.00,
	// 	CategoryID: 1,
	// }
	// db.Create(&productNotebook)

	// productMouse := Product{
	// 	Name:       "Mouse",
	// 	Price:      300.00,
	// 	CategoryID: 1,
	// }
	// productMonitor := Product{
	// 	Name:       "Mouse",
	// 	Price:      700.00,
	// 	CategoryID: 1,
	// }

	// db.Create(&productNotebook)
	// db.Create(&productMouse)
	// db.Create(&productMonitor)

	// serialNotebook := SerialNumber{
	// 	Number:    "123456",
	// 	ProductID: 1,
	// }
	// serialMouse := SerialNumber{
	// 	Number:    "1234567",
	// 	ProductID: 2,
	// }
	// serialMonitor := SerialNumber{
	// 	Number:    "12345678",
	// 	ProductID: 3,
	// }
	// db.Create(&serialNotebook)
	// db.Create(&serialMouse)
	// db.Create(&serialMonitor)

	// * create com relacionamento "um pra muitos"

	// * select one
	// var product Product
	// var productMouse Product

	// db.First(&product, 2)
	// fmt.Println(product)
	// fmt.Println("\n")

	// db.First(&productMouse, "name = ?", "Mouse")
	// fmt.Println(productMouse)
	// fmt.Println("\n")

	// * select all
	// var products []Product
	// db.Find(&products)

	// for _, item := range products {
	// 	fmt.Println(item)
	// }
	// fmt.Println("\n")

	// * select com relacionamento um pra um
	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	// * select com relacionamento um pra muitos
	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		for _, productItem := range category.Products {
			println(productItem.Name, productItem.SerialNumber.Number, category.Name)
			println("\n")
		}
	}

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
	// var products []Product
	// db.Where("price >= ?", 1000).Find(&products)
	// for _, item := range products {
	// 	fmt.Println(item)
	// }

	// var products []Product
	// db.Where("name like ?", "%book%").Where("price > ?", 2500).Find(&products)
	// for _, item := range products {
	// 	fmt.Println(item)
	// }

	// * update
	// var pUpdate Product
	// db.First(&pUpdate, 1)
	// pUpdate.Name = "update name Notebook"
	// db.Save(&pUpdate)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2)

	// * Delete
	// var pMouse Product
	// db.First(&pMouse, "name like ?", "%notebook%")
	// fmt.Println(pMouse)
	// db.Delete(&pMouse)

}
