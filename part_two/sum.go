package parttwo

import (
	"strconv"
)

// Add ...
func Add(v1 int, v2 int) int {
	return v1 + v2
}

// FloatToString ...
func FloatToString(f1 float64, f2 float64) string {
	return strconv.FormatFloat(f1/f2*100, 'f', 2, 32) + "%"
}
