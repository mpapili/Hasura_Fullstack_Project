package models

import (
	"gorm.io/gorm"
)

type Meetup struct {
	gorm.Model
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Address     string `json:"address"`
	Image       string `json:"image"`
	Description string `json:"description"`
}

func (Meetup) TableName() string {
	return "meetups_meetup"
}
