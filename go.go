package isodd

import (
	"fmt"
	"strconv"
)

// IsOdd returns true if the number is odd WITHOUT USING THE %
func IsOdd(n int) bool {
	var num float64 = float64(n) / 2

	n1, _ := strconv.ParseFloat(fmt.Sprintf("%f", num), 64)
	n2, _ := strconv.ParseFloat(fmt.Sprintf("%.0f", num), 64)

	if n1 != n2 {
		return true
	}
	return false
}
