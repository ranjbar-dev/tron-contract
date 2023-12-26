package main

import (
	"strconv"
)

func convertToInt(value interface{}) int {
	switch v := value.(type) {
	case int:
		return v
	case float64:
		return int(v)
	case float32:
		return int(v)
	case string:
		temp, _ := strconv.Atoi(v)
		return temp
	default:
		return 0
	}
}
