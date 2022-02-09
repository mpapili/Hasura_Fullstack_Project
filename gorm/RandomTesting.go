/*
package main // set to main once done

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"fmt"

	"gormtest/models"
)

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	check_err(err)

	// Migrate the Schema
	db.AutoMigrate(&models.Product{})

	// Create a record for our product table
	db.Create(&models.Product{Code: "D42", Price: 100})

	// Read
	var product models.Product // we will load into this
	// get first product with code 42
	db.First(&product, "code = ?", "D42")
	fmt.Println(product.Price)

	// Update Product's Price to 200
	db.Model(&product).Update("Price", 200)
	// Update multiple fields on our singular record
	db.Model(&product).Updates(models.Product{Price: 200, Code: "F42"})
	// or do the same thing with a map object!
	mymap := map[string]interface{}{"Price": 300, "Code": "F42"}
	db.Model(&product).Updates(mymap)

	// can we find multiples?
	var products []models.Product
	db.Find(&products, "price = ?", 300)
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i].Code)
	}

	// Delete the products
	db.Delete(&products)
	// NOTE - this does a gorm SOFT DELETE which updates "deleted_at"
	// if you truly want to nuke them do:
	db.Find(&products, "price > ?", -1) // get all records without WHERE clauses
	db.Unscoped().Delete(&products)

}
*/