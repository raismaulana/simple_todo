package dto

import "database/sql"

type UserUpdateDTO struct {
	Name     string         `json:"name" form:"name" binding:"required"`
	Address  string         `json:"address" form:"address" binding:"required"`
	Photo    sql.NullString `json:"photo,omitempty" form:"photo,omitempty"`
	Password string         `json:"password,omitempty" form:"password,omitempty"`
	Role     string         `json:"role,omitempty" form:"role,omitempty"`
}
