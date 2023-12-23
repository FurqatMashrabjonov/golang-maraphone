package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		return
	}

	defer dir.Close()

	files, err := dir.Readdir(-1)

	if err != nil {
		return
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

}
