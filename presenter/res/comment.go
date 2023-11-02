package res

import "time"

type Comment struct {
	ID          string     `json:"commentId"`
	ThreadID    string     `json:"threadId"`
	File        *File      `json:"file"`
	TextContent string     `json:"textContent"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}
