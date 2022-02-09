package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gormtest/models"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	models.DB.Find(&products)

	c.JSON(http.StatusOK, gin.H{"data": products})
}
