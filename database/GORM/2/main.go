package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       uint `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:category_products;"`
}

type Product struct {
	ID           uint `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   uint
	Categories   []Category `gorm:"many2many:category_products;"`
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        uint `gorm:"primaryKey"`
	Number    string
	ProductID uint
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	category := Category{
		Name: "Cozinha",
	}
	createCategory(db, &category)

	category2 := Category{
		Name: "Eletronico",
	}
	createCategory(db, &category2)

	product := Product{
		Name:       "Air fryer",
		Price:      2300,
		Categories: []Category{category, category2},
	}
	createProduct(db, &product)

	//db.Create(&SerialNumber{
	//	Number:    "085232",
	//	ProductID: product.ID,
	//})

	//var products []Product
	//db.Preload("Category").Preload("SerialNumber").Find(&products)
	//for _, product := range products {
	//	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	//}

	var categories []Category

	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, category := range categories {
		fmt.Println(category.Name)
		for _, product := range category.Products {
			fmt.Println("- ", product.Name)
		}
	}

}

func createCategory(db *gorm.DB, category *Category) {
	db.Create(&category)
}

func createProduct(db *gorm.DB, product *Product) {
	db.Create(&product)
}
