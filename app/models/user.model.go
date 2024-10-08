package models

import (
	"app/config"
	"time"
)

type User struct {
	ID            uint      `json:"id" gorm:"primaryKey"`
	Name          string    `json:"name" gorm:"type:varchar(100)"`
	Email         string    `json:"email" gorm:"unique;type:varchar(100)"`
	Password_hash string    `json:"-" gorm:"type:varchar(255)"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Register struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoggedIn struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func GetCustomerByEmail(email string) User {
	var customer User
	if err := config.DB.Where("email = ?", email).First(&customer).Error; err != nil {
		return customer
	}

	return customer
}

func CreateCustomer(customer *User) error {
	return config.DB.Create(customer).Error
}
