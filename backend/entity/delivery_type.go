package entity

import (
	"gorm.io/gorm"
)

type DeliveryType struct {
	gorm.Model
	Type    string
	Payment []Payment `gorm:"foreignKey:DeliveryTypeID"`
}