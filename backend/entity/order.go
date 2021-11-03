package entity

import (
	"time"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	PreorderID int 
	Statusorder string
	StatusID int
	OwnerID    *uint
	Ordertime  time.Time
	Owner      User		`gorm:"references:id"`
	Payment    []Payment `gorm:"foreignKey:OrderID"`
	Returns    []Return `gorm:"foreignKey:OrderID"`
}