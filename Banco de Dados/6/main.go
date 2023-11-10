package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product `gorm:"many2many:category_products;"`
}

type Product struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:category_products;"`
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{})

	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// category2 := Category{Name: "Eletronicos"}
	// db.Create(&category2)

	// category3 := Category{Name: "Utensilios"}
	// db.Create(&category3)

	// db.Create(&Product{
	// 	Name:       "Panela de ferro",
	// 	Price:      2000,
	// 	Categories: []Category{category, category2},
	// })

	// db.Create(&Product{
	// 	Name:       "Termometro",
	// 	Price:      199,
	// 	Categories: []Category{category, category2, category3},
	// })

	// db.Create(&Product{
	// 	Name:       "Espatula",
	// 	Price:      20,
	// 	Categories: []Category{category, category3},
	// })

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
