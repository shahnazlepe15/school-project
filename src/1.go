package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42)
	result := rand.Intn(100)
	fmt.Println("The random number is:", result)
}
