package main

import (
	"fmt"
)

func main() {
	city := PopcenterNew()
	fmt.Printf(city.String())

	/*for i := 0; i < 100; i++ {
		fmt.Printf("%+v\n", PersonNew())
	}*/
}

