package models

import "time"

type UserRole string

const (
	Customer = UserRole("customer")
	Coach    = UserRole("coach")
)

type User struct {
	ID     int
	Name   string   `gorm:"not null"`
	Email  string   `gorm:"unique"`
	Role   UserRole `gorm:"not null"`
	PicURL *string
	DOB    *time.Time
}
