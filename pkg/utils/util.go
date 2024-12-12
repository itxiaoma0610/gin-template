package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func RandomNumeric(codeLength int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if codeLength <= 0 {
		panic("code length " + strconv.Itoa(codeLength) + "must be greater than 0")
	}
	value := ""
	for index := 0; index < codeLength; index++ {
		value += strconv.Itoa(r.Intn(9))
	}
	return value
}
