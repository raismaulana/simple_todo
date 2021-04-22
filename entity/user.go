package entity

import (
	"database/sql"
	"time"
)

type User struct {
	ID              uint64         `gorm:"primary_key:auto_increment" json:"id"`
	Name            string         `gorm:"type:varchar(255) not null" json:"name"`
	Address         string         `gorm:"type:text not null" json:"address"`
	Photo           sql.NullString `gorm:"type:text" json:"photo"`
	Email           string         `gorm:"type:varchar(255) not null unique" json:"email"`
	EmailVerifiedAt sql.NullTime   `gorm:"type:timestamp NULL;default:NULL" json:"-"`
	Password        string         `gorm:"type:text not null" json:"-"`
	Role            string         `gorm:"type:enum('user', 'admin') not null;default:'user'" json:"-"`
	Token           string         `gorm:"-" json:"token,omitempty"`
	DeletedAt       sql.NullTime   `gorm:"type:timestamp NULL;default:NULL" json:"-"`
	CreatedAt       time.Time      `gorm:"type:timestamp not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt       time.Time      `gorm:"type:timestamp not null;default:CURRENT_TIMESTAMP" json:"-"`
	Loan            []Loan         `json:"loan,omitempty"`
}

func (User) TableName() string {
	return "User"
}
