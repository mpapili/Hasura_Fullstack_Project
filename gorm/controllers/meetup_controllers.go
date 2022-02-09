package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gormtest/models"
)

func GetMeetups(c *gin.Context) {
	var meetups []models.Meetup
	models.DB.Find(&meetups)

	c.JSON(http.StatusOK, gin.H{"data": meetups})
}
