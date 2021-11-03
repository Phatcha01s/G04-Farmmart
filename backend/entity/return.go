package entity

import (
	"time"
	"gorm.io/gorm"
)
type Return struct {
	gorm.Model
	// OwnerID ทำหน้าที่เป็น FK
	OwnerID *uint
	Owner   User `gorm:"references:id"`

	// OderID ทำหน้าที่เป็น FK
	OrderID *uint
	Order   Order `gorm:"references:id"`

	// StaffID ทำหน้าที่เป็น FK
	StaffID *uint
	Staff   Staff `gorm:"references:id"`

	Reason     string
	Returndate time.Time
}