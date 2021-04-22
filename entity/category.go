package entity

import (
	"database/sql"
	"time"
)

type Category struct {
	Id        uint64       `gorm:"primary_key:auto_increment"`
	Name      string       `gorm:"type:varchar(255) not null"`
	DeletedAt sql.NullTime `gorm:"type:timestamp NULL;default:NULL"`
	CreatedAt time.Time    `gorm:"type:timestamp not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time    `gorm:"type:timestamp not null;default:CURRENT_TIMESTAMP"`
	Book      []Book
}

func (Category) TableName() string {
	return "Category"
}
