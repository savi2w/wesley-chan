package model

import "time"

type Thread struct {
	ID          string     `db:"thread_id"`
	BoardID     string     `db:"board_id"`
	FileID      *string    `db:"file_id"`
	TextContent string     `db:"text_content"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
