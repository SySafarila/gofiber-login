package models

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type User struct {
	Id          string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	IncrementId int       `gorm:"uniqueIndex;autoIncrement" json:"increment_id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Username    *string   `gorm:"type:varchar(255);uniqueIndex" json:"username,omitempty"`
	Email       string    `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Password    string    `gorm:"type:varchar(255);not null" json:"password"`
	IsVerified  bool      `gorm:"default:false" json:"is_verified"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	//DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"` // Opsional: untuk soft delete
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

func (User) TableName() string {
	return "users"
}

type TokenClaims struct {
	UserId string `json:"user_id"`
	Iat    int64
	Exp    int64
	jwt.RegisteredClaims
}

func (u User) CreateToken() string {
	var key []byte
	var token *jwt.Token
	var signedToken string

	claims := jwt.MapClaims{
		"user_id": u.Id,
		"iat":     time.Now().Unix(),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	key = []byte("MYSECRETKEY")
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ = token.SignedString(key)

	// save token to redis for a while

	// save token to redis for a while

	return signedToken
}

func (u User) CreateUserResponse() UserResponse {
	return UserResponse{
		Id:        u.Id,
		Name:      u.Name,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
