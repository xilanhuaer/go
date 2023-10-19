package utils

import "strconv"

// limit offset err
func PageUtil(page, page_size string) (int, int, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 0 {
		return 0, 0, err
	}
	limit, err := strconv.Atoi(page_size)
	if err != nil || limit < 0 {
		return 0, 0, err
	}
	offset := (pageInt - 1) * limit
	return limit, offset, nil
}
