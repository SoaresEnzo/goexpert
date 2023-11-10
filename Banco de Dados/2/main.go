package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:"primary_key"`
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
	// db.Create(&Product{Name: "Laptop", Price: 1000})

	// products := []Product{
	// 	{Name: "Laptop", Price: 1000},
	// 	{Name: "Mouse", Price: 100},
	// 	{Name: "Keyboard", Price: 200},
	// }

	// db.Create(&products)

	// var product Product
	// db.First(&product, 1)
	// fmt.Println(product)

	// db.First(&product, "name = ?", "Mouse")
	// fmt.Println(product)

	// var products []Product
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Where("price > ?", 900).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var products []Product
	// db.Where("name LIKE ?", "%top%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// var p Product
	// db.First(&p, 1)
	// p.Name = "New mouse"
	// db.Save(&p)

	// var p2 Product
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)

	// db.Delete(&p2)
}
