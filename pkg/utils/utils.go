package utils

import "strconv"

func GetInt(a interface{}) int {
	switch a := a.(type) {
	case int:
		return a
	case int64:
		return int(a)
	case float64:
		return int(a)
	case string:
		val, err := strconv.Atoi(a)
		if err != nil {
			return 0
		}
		return val
	default:
		return 0
	}
}
