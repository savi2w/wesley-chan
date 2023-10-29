package payload

import (
	"errors"
	"strconv"
)

func GetOffsetByPage(page string, itemsPerPage int64) (int64, error) {
	if len(page) <= 0 {
		return 0, nil
	}

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return 0, errors.New("page should be a number")
	}

	return pageInt * itemsPerPage, nil
}
