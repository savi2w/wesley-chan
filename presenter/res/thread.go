package res

import "time"

type Thread struct {
	ID          string     `json:"threadId"`
	BoardID     string     `json:"boardId"`
	FileID      *string    `json:"fileId"`
	Subject     string     `json:"subject"`
	TextContent string     `json:"textContent"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}
