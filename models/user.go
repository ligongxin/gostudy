package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID         int64          `json:"id" db:"id"`
	UserId     int64          `json:"user_id" db:"user_id"`
	Username   string         `json:"username" db:"username"`
	Password   string         `json:"password" db:"password"`
	Email      sql.NullString `json:"email" db:"email"`
	Gender     int8           `json:"gender" db:"gender"`
	CreateTime time.Time      `json:"create_time" db:"create_time"`
	UpdateTime time.Time      `json:"update_time" db:"update_time"`
}
