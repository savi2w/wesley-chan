package model

import "time"

type File struct {
	ID        int64      `db:"file_id"`
	Path      string     `db:"path"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
