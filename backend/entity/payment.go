package entity

import (
	"time"
	"gorm.io/gorm"
)

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
