package utils

import (
	"math/rand"
	"strconv"
)

func RandomID() string {
	return strconv.Itoa(int(rand.Uint32()))
}
