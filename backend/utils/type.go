package utils

import "strconv"

func StringToInt64(s string) int64 {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int64(n)
}
