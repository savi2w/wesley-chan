package payload

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/savi2w/wesley-chan/presenter/req"
)

func GetComment(rc io.ReadCloser) (r *req.Comment, err error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("invalid read closer")
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, errors.New("invalid json payload")
	}

	if len(r.ThreadID) != 36 {
		return nil, errors.New("invalid thread uuid")
	}

	if r.FileID != nil {
		length := len(*r.FileID)

		if length < 38 && length > 56 {
			return nil, errors.New("invalid file uuid")
		}
	}

	if len(r.TextContent) <= 0 {
		return nil, errors.New("you cannot send an empty comment, please type something")
	}

	if len(r.TextContent) > 2048 {
		return nil, errors.New("comment too long, please keep it at under 2048 characters")
	}

	return r, nil
}
