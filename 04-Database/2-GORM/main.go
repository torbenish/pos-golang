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
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{})

	//create
	// db.Create(&Product{
	// 	Name:  "Laptop",
	// 	Price: 1000.00,
	// })

	// create batch
	// products := []Product{
	// 	{Name: "Laptop", Price: 1000.00},
	// 	{Name: "Mouse", Price: 50.00},
	// 	{Name: "Keyboard", Price: 75.00},
	// }
	// db.Create(&products)

	// select one
	// var product Product
	// // db.First(&product, 3)
	// // fmt.Println(product)
	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// select all
	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// where
	// var products []Product
	// db.Where("price > ?", 90).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Where("name LIKE ?", "%top").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// update
	// var p Product
	// db.First(&p, 1)
	// p.Name = "Nem Mouse"
	// db.Save(&p)

	// delete
	var p2 Product
	db.First(&p2, 1)
	fmt.Println(p2.Name)
	db.Delete(&p2)
}
