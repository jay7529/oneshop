package utils

import (
	"strconv"
)

// string to int
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

// int to string
func IntToString(i int) string {
	s := strconv.Itoa(i)

	return s
}

// int64 to string
func Int64ToString(i int64) string {
	s := strconv.FormatInt(i, 10)

	return s
}
