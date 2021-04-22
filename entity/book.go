package entity

import (
	"database/sql"
	"time"
)

type Book struct {
	Id          uint64 `gorm:"primary_key:auto_increment"`
	Title       string `gorm:"type:text not null"`
	Description string `gorm:"type:text not null"`
	Author      string `gorm:"type:text not null"`
	Year        string `gorm:"type:year(4) not null"`
	CategoryID  uint64
	Stock       uint         `gorm:"type:int(11) not null"`
	Status      string       `gorm:"type:enum('tersedia', 'kosong') not null;default:'kosong'"`
	DeletedAt   sql.NullTime `gorm:"type:timestamp NULL;default:NULL"`
	CreatedAt   time.Time    `gorm:"type:timestamp not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time    `gorm:"type:timestamp not null;default:CURRENT_TIMESTAMP"`
	Loan        []Loan
}

func (Book) TableName() string {
	return "Book"
}
