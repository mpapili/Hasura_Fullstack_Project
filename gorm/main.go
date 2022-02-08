package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

	"fmt"
	"net/http"

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
	for i:=0; i<len(products); i++ {
		fmt.Println(products[i].Code)
	}

	// set up a gin router
	r := gin.Default()

	// url mapping for "/"
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello, world!"})
	})

	// url mapping to fetch products
	r.GET("/products", func(c *gin.Context) {
		// fetch our products
		var prods []models.Product
		db.Find(&prods)

		// return then in JSON format
		c.JSON(http.StatusOK, gin.H{"data": prods})
	})

	// fire up the server hardcoded to 8081
	r.Run(":8081")

	// Delete the products
	db.Delete(&products)
	// NOTE - this does a gorm SOFT DELETE which updates "deleted_at"
	// if you truly want to nuke them do:
	db.Find(&products) // get all records without WHERE clauses
	db.Unscoped().Delete(&products)

}
