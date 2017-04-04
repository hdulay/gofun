package main

import (
	"fmt"
)

func Sqrt(s float64) float64 {

	x := float64(8 * 10 * 10)

	for {
		prev := x
		x = (x + s/x) / 2
		fmt.Println(x);
		if prev == x { break }
	}

	return x;
}

func main() {
	fmt.Println(Sqrt(2))
}
