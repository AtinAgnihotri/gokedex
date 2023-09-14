package utils

import "math/rand"

// var randSeed  := rand.New(rand.NewSource(99))

var randSeed *rand.Rand

func GetRandom(upperBound int) int {
	if randSeed == nil {
		randSeed = rand.New(rand.NewSource(99))
	}
	return randSeed.Intn(upperBound)
}

func CheckIfExists[T comparable](arr []T, val T) bool {
	exists := false
	for _, item := range arr {
		if item == val {
			exists = true
		}
	}
	return exists
}
