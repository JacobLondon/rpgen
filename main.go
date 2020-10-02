package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	town := Town_new()
	fmt.Printf("%+v\n", town)
	town.Write()
}
