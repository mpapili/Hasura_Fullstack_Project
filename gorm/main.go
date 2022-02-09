package main

import (
	"github.com/gin-gonic/gin"

	"net/http"

	"gormtest/controllers"
	"gormtest/models"
)

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	// connect our controllers to the db
	models.ConnectDatabase()

	// set up a gin router
	r := gin.Default()
	r.Use(CORSMiddleware())

	// url mapping for "/"
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello, world!"})
	})

	// url mapping to fetch all meetups
	r.GET("/meetups", func(c *gin.Context) {
		// fetch our products
		controllers.GetMeetups(c)
	})

	// url mapping to fetch products
	r.GET("/products", func(c *gin.Context) {
		// fetch our products
		controllers.GetProducts(c)
	})

	// fire up the server hardcoded to 8081
	r.Run(":8081")

}
