package main

import "fmt"

type Age int

func (a Age) printValue() {
	fmt.Println("Value is: ", a)
}

func main() {
	a := Age(23)
	a.printValue()
}
