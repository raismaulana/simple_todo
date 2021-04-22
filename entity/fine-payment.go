package entity

import (
	"database/sql"
	"time"
)

type FinePayment struct {
	Id        uint64       `gorm:"primary_key:auto_increment"`
	Receipt   string       `gorm:"type:varchar(225) NULL"`
	Amount    float64      `gorm:"type:double not null"`
	LoanID    uint64       `gorm:"not null"`
	Loan      Loan         `gorm:"foreignkey:LoanID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeletedAt sql.NullTime `gorm:"type:timestamp NULL;default:NULL"`
	CreatedAt time.Time    `gorm:"type:timestamp not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time    `gorm:"type:timestamp not null;default:CURRENT_TIMESTAMP"`
}

func (FinePayment) TableName() string {
	return "FinePayment"
}
