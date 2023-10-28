package model

import "time"

type Board struct {
	ID          int64      `db:"board_id"`
	Name        string     `db:"name"`
	Slug        string     `db:"slug"`
	Description string     `db:"description"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
