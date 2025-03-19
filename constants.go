package main

import (
	"fmt"
	"math"
)

const a string = "Furqat"

func main() {
	fmt.Println(a)

	const b = 5000000
	const d = 3e20 / b

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(b))
}
