package entity

import (
	"database/sql"
	"time"
)

type Loan struct {
	Id           uint64 `gorm:"primary_key:auto_increment"`
	UserID       uint64
	BookID       uint64
	BorrowedDate time.Time    `gorm:"type:date not null"`
	DueDate      time.Time    `gorm:"type:date not null"`
	ReturnDate   time.Time    `gorm:"type:date NULL"`
	DeletedAt    sql.NullTime `gorm:"type:timestamp NULL;default:NULL"`
	CreatedAt    time.Time    `gorm:"type:timestamp not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time    `gorm:"type:timestamp not null;default:CURRENT_TIMESTAMP"`
}

func (Loan) TableName() string {
	return "Loan"
}
