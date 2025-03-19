package main

import (
	"fmt"
)

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	fmt.Println("---------------------")

	for i := 1; i <= 5; i += 1 {
		fmt.Println(i)
	}

	fmt.Println("---------------------")

	for i := range 3 {
		fmt.Println(i)
	}

	// for {
	// 	fmt.Println(i)
	// }

	fmt.Println("---------------------")

	for i := range 6 {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}

}
