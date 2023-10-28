package model

import "time"

type Comment struct {
	ID          int64      `db:"comment_id"`
	ThreadID    int64      `db:"thread_id"`
	FileID      *int64     `db:"file_id"`
	TextContent string     `db:"text_content"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}
