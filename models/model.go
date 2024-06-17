package models

import (
	"time"
)

type Customer struct {
	ID      uint    `gorm:"primaryKey"`
    Name 	string `json:"name" gorm:"size:255"`
    Code 	string `json:"code" gorm:"size:100;unique"`
}

type Order struct {
	ID        uint      `gorm:"primaryKey"`
    Item      string    `gorm:"size:255"`
    Amount    float64	`json:"amount"`
	Time      time.Time
    CustomerID uint		`json:"customer_id"`
}