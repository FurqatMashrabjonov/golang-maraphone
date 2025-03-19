package main

import "fmt"

func main() {
	if 7%2 == 1 {
		fmt.Println("7 is odd number")
	}

	if num := 9; num < 10 {
		fmt.Println("9 is less than 10")
	}
}
