package entity

import (
	"time"
	// "gorm.io/gorm"
	// "gorm.io/gorm"
)

type User struct {
	ID         int    `json:"id"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Email      string `json:"email"`
	Birthday   time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Password   string `json:"password"`
	Age        string `json:"age"`
	Token string `json:"token"`
	UUID string `json:"uuid"`
}
