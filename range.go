package main

import "fmt"

func main() {
	nums := []int{3, 4, 5}

	for i, num := range nums {
		fmt.Println("index: ", i, "num: ", num)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
