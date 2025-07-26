package utils

import (
	"strconv"

	"github.com/charmbracelet/huh"
)

// credit to "github.com/Broderick-Westrope/charmutils" for this form range function

// HuhIntRangeOptions creates a new slice of integer options using the lower and upper bound (inclusive).
func HuhIntRangeOptions(lower, upper int) []huh.Option[int] {
	opts := make([]huh.Option[int], (upper-lower)+1)
	for i := range opts {
		v := i + lower
		opts[i] = huh.Option[int]{Key: strconv.Itoa(v), Value: v}
	}
	return opts
}
