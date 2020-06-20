package main

import (
	"math/rand"
)

func pick_one(slice []string) string {
	return slice[rand.Intn(len(slice))]
}
