package req

type Comment struct {
	ThreadID    string  `json:"threadId"`
	FileID      *string `json:"fileId"`
	TextContent string  `json:"textContent"`
}
