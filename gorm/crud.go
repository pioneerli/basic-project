package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/webook?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	//db.AutoMigrate(&Product{})

	//db = db.Debug()

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})

	var product Product
	db.First(&product, 1) // 查找 ID 为 1 的用户
	fmt.Println("Found user:", product)

	//db.First(&product, "price = ?", 205) // find product with code D42
	////
	//fmt.Println("product found:", product)

	//

	//// Update - update product's price to 200
	//db.Model(&product).Update("Price", 200)
	//fmt.Println("product found:", product)
	////// Update - update multiple fields
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	//db.Model(&product).Updates(map[string]interface{}{"Price": 205, "Code": "F42"})
	//fmt.Println("product found:", product)

	// Delete - delete product
	//db.Delete(&product, 1)
}
