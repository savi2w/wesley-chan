package payload

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/savi2w/wesley-chan/presenter/req"
)

func GetThread(rc io.ReadCloser) (r *req.Thread, err error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("invalid read closer")
	}

	if err := json.Unmarshal(body, r); err != nil {
		return nil, errors.New("invalid json payload")
	}

	if len(r.BoardID) != 36 {
		return nil, errors.New("invalid board uuid")
	}

	if r.FileID != nil {
		length := len(*r.FileID)

		if length < 38 && length > 56 {
			return nil, errors.New("invalid file uuid")
		}
	}

	if len(r.Subject) <= 0 {
		return nil, errors.New("you cannot send an empty subject, please type something")
	}

	if len(r.Subject) > 256 {
		return nil, errors.New("subject too long, please keep it at under 256 characters")
	}

	if len(r.TextContent) <= 0 {
		return nil, errors.New("you cannot send an empty thread, please type something")
	}

	if len(r.TextContent) > 2048 {
		return nil, errors.New("thread too long, please keep it at under 2048 characters")
	}

	return r, nil
}
