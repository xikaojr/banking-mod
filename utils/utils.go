// utils/utils.go

package utils

import "math/rand"

var lastAccountNumber int = 1000

func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func NextAccountNumber() int {
	lastAccountNumber++
	return lastAccountNumber
}
