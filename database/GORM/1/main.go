package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

type Category struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{})

	////Cria um produto
	db.Create(&Product{
		Name:  "Maquina de lavar",
		Price: 2030.00,
	})
	//
	////Cria varios produtos
	//produts := []Product{
	//	{Name: "Notebook", Price: 5320.00},
	//	{Name: "Maquina de lavar", Price: 2030.00},
	//	{Name: "Camera", Price: 4252.00},
	//}
	//db.Create(&produts)
	//
	////seleciona uma linha
	//var product Product
	//db.First(&product, 1)
	//db.First(&product, "name = ?", "Maquina de lavar")
	//
	////seleciona tudo
	////var products []Product
	////db.Find(&products)
	////for _, product := range products {
	////	fmt.Println(product)
	////}
	////
	////var products []Product
	////db.Where("where > ?", "").Find(&products)
	////for _, product := range products {
	////	fmt.Println(product)
	////}
	//
	//// Alterando
	var p Product
	db.First(&p, 1)
	p.Name = "Notebook"
	db.Save(&p)
	////
	////Excluindo
	var p2 Product
	db.First(&p2, 1)
	db.Delete(&p2)
}
