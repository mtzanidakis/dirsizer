package util

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// ByteCountIEC returns a human-readable byte string of the form 10M, 12.5K, and so forth
func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c",
		float64(b)/float64(div), "KMGTPE"[exp])
}

// IECToBytes converts a string like "10M" into an integer representing the number of bytes.
func IECToBytes(iec string) (int64, error) {
	const unit = 1024
	exp := map[string]float64{"K": 1, "M": 2, "G": 3, "T": 4, "P": 5, "E": 6}

	num, err := strconv.ParseInt(iec[:len(iec)-1], 10, 64)
	if err != nil {
		return 0, err
	}

	b := num * int64(math.Pow(unit, exp[strings.ToUpper(iec[len(iec)-1:])]))

	return b, nil
}
