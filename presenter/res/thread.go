package res

type Thread struct {
	BoardID     string  `json:"boardId"`
	FileID      *string `json:"fileId"`
	Subject     string  `json:"subject"`
	TextContent string  `json:"textContent"`
}
