package models

import (
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID         int    `json:"id" gorm:"column:id;"`
	UserName   string `json:"user_name" gorm:"column:user_name;"`
	Email      string `json:"email" gorm:"column:email;"`
	Password   string `json:"password" gorm:"column:password;size:1000"`
	Address    string `json:"address" gorm:"column:address;"`
	PrivateKey string `json:"private_key" gorm:"column:private_key;"`
	Role       string `json:"role" gorm:"column:role;"`
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func Santize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
