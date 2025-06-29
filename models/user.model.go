package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id        string         `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	Username  sql.NullString `json:"username"`
	CreatedAt time.Time      `json:"created_at" gorm:"index"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type UserResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	// Password  string    `json:"password"`
	Username  *string   `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
