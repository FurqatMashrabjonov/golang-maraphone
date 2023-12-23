package main

import (
	"fmt"
	"strings"
)

func main() {

	str := "a-b-c-d-f"
	var arr []string
	arr = strings.Split(str, "-")

	fmt.Println(arr[2:4])

}
