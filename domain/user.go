package domain

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type User struct {
	ID              ulid.ULID `json:"id" gorm:"type:uuid;not null;primary key"`
	Name            string    `json:"name" gorm:"type:varchar(50);not null"`
	Email           string    `json:"email" gorm:"type:varchar(20); not null;unique"`
	Password        string    `json:"password" gorm:"type:varchar(20); not null"`
	PhoneNumber     string    `json:"phone_number" gorm:"type:varchar(15); not null"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	EmailVerifiedAt time.Time `json:"deleted_at"`
}
