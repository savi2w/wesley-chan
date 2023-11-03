package payload

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"

	"github.com/savi2w/wesley-chan/presenter/req"
)

func GetBoard(rc io.ReadCloser) (r *req.Board, err error) {
	defer rc.Close()

	body, err := io.ReadAll(rc)
	if err != nil {
		return nil, errors.New("invalid read closer")
	}

	if err := json.Unmarshal(body, &r); err != nil {
		return nil, errors.New("invalid json payload")
	}

	if len(r.Name) <= 0 {
		return nil, errors.New("you cannot send an empty name, please type something")
	}

	if len(r.Name) > 64 {
		return nil, errors.New("name too long, please keep it at under 64 characters")
	}

	if len(r.Slug) <= 0 {
		return nil, errors.New("you cannot send an empty slug, please type something")
	}

	if len(r.Slug) > 16 {
		return nil, errors.New("slug too long, please keep it at under 16 characters")
	}

	if len(r.Description) <= 0 {
		return nil, errors.New("you cannot send an empty description, please type something")
	}

	if len(r.Description) > 256 {
		return nil, errors.New("description too long, please keep it at under 256 characters")
	}

	return r, nil
}

func GetSlug(slug string) (string, error) {
	if len(slug) <= 0 {
		return "", errors.New("you cannot send an empty slug, please type something")
	}

	if len(slug) > 16 {
		return "", errors.New("slug too long, please keep it at under 16 characters")
	}

	// Check if slug has any non-alphanumeric characters using RegExp
	if !regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(slug) {
		return "", errors.New("slug can only contain alphanumeric characters")
	}

	return slug, nil
}
