package helpers

import "strconv"

func StringToInt64(s string) (*int64, error) {
	num, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, err
	}
	return &num, nil
}
