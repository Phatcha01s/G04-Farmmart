package entity

import (
	"gorm.io/gorm"
)
type Staff struct {
	gorm.Model
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Password string
	ProductStocks []ProductStock `gorm:"foreignKey:StaffID"`
	Returns []Return `gorm:"foreignKey:StaffID"`
}