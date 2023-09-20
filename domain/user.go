package domain

import (
	"time"
)

type User struct {
	ID          string    `json:"id" gorm:"type:varchar(100);not null;primary key"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Email       string    `json:"email" gorm:"type:varchar(100); not null;unique"`
	Password    string    `json:"password" gorm:"type:varchar(100); not null"`
	PhoneNumber string    `json:"phone_number" gorm:"type:varchar(15); not null"`
	OTPCode     string    `json:"otp_code" gorm:"type:char(4)"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IsVerified  bool      `json:"is_verified" gorm:"type:boolean;default:false"`
}

type RegistBind struct {
	Name                  string `json:"name" binding:"required,max=20,min=3"`
	PhoneNumber           string `json:"phone_number" binding:"required,numeric,min=10"`
	Email                 string `json:"email" binding:"required,email"`
	Password              string `json:"password" binding:"required,min=8"`
	Verification_Password string `json:"verification_password" binding:"required,min=8"`
}

type VerifyAccountBind struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required,max=4,min=4"`
}
