package models

import (
	"time"
)

type Order struct {
	OrderID      uint   `gorm:"primaryKey"`
	CustomerName string `gorm:"not null;type:varchar(191)"`
	Items        []Item
	OrderedAt    time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Item struct {
	ItemID      uint   `gorm:"primaryKey"`
	ItemCode    string `gorm:"not null;type:varchar(191)"`
	Description string `gorm:"not null;type:text"`
	Quantity    uint   `gorm:"not null;type:int"`
	OrderID     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
