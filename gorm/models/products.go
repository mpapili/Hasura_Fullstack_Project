package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code	string	`json:"code"`
	Price	int	`json:"price"`
	Name	string	`json:"name"`
}

