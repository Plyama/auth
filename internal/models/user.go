package models

type UserRole string

const (
	Customer = UserRole("customer")
	Coach    = UserRole("coach")
)

type User struct {
	ID    int32
	Name  string   `gorm:"not null"`
	Email string   `gorm:"unique"`
	Role  UserRole `gorm:"not null"`
}
