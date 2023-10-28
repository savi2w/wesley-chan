package model

import "time"

type Thread struct {
	ID          int64      `db:"thread_id"`
	BoardID     int64      `db:"board_id"`
	FileID      *int64     `db:"file_id"`
	TextContent string     `db:"text_content"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
