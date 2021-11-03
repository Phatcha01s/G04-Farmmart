package entity

import (
	"time"
	"gorm.io/gorm"
)

type PaymentMethod struct {
	gorm.Model
	Method  string
	Payment []Payment `gorm:"foreignKey:PaymentMethodID"`
}

type DeliveryType struct {
	gorm.Model
	Type    string
	Payment []Payment `gorm:"foreignKey:DeliveryTypeID"`
}

type Payment struct {
	gorm.Model
	Phone	string
	PaymentTime time.Time

	OrderID *uint
	Order   Order	`gorm:"references:id"`

	PaymentMethodID *uint
	PaymentMethod   PaymentMethod 	`gorm:"references:id"`

	DeliveryTypeID *uint
	DeliveryType   DeliveryType	`gorm:"references:id"`
}
