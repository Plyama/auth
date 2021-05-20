package models

type User struct {
	ID    int32
	Name  string
	Email string `gorm:"unique"`
	Token string
}
