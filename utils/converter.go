package utils

import (
	"fmt"
	"strconv"
)

// HexToInt converts a hexadecimal string to an integer.
func HexToInt(hexStr string) (int, error) {
	// Remove the "0x" prefix if it exists
	if len(hexStr) > 2 && hexStr[:2] == "0x" {
		hexStr = hexStr[2:]
	}

	// Convert the hexadecimal string to an integer
	value, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid hexadecimal string: %s", hexStr)
	}

	return int(value), nil
}
