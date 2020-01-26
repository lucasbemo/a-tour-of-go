package main

import (
	"fmt"
	"math"
)

func sqrt(x  float64) string {
	if x < 0 {
		return sqrt(-x) + "1"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}