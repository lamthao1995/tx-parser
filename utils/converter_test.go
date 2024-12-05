package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHexToInt(t *testing.T) {
	hexStr := "0x4b7"
	result, err := HexToInt(hexStr)
	assert.NoError(t, err, "Expected no error for valid hex string")
	assert.Equal(t, result, 1207, "Expected integer result to be 1207")

	hexStrInvalid := "0xGHI"
	_, err = HexToInt(hexStrInvalid)
	assert.Error(t, err, "Expected error for invalid hex string")
}
