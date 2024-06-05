package utils

import "time"

// ParseTime 解析指定格式的时间
func ParseTime(timeStr string, format string) time.Time {
	t, err := time.ParseInLocation(format, timeStr, time.Local)
	if err != nil {
		return time.Now()
	}
	return t
}
