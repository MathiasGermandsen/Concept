package models

import (
	"time"

	"gorm.io/gorm"
)

type Inquiry struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	CustomerName   string         `json:"customer_name" gorm:"not null"`
	Email          string         `json:"email" gorm:"not null"`
	PhoneNumber    string         `json:"phone_number"`
	EventDate      string         `json:"event_date"`
	EventLocation  string         `json:"event_location"`
	Message        string         `json:"message" gorm:"type:text"`
	EstimatedPrice float64        `json:"estimated_price"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
