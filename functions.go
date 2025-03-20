package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(324, 324))

	_, b := vals()

	fmt.Println(b)

	nums := []int{31, 3, 14, 12, 4, 123, 123, 12, 3}
	fmt.Println("sum:", sum(nums...))

	nextVal := initSqn()

	fmt.Println(nextVal())
	fmt.Println(nextVal())
	fmt.Println(nextVal())
	fmt.Println(nextVal())

	fmt.Println("4 factorial: ", factorial(7))

	fmt.Println("fib: ", fib(7))
}

func vals() (int, int) {
	return 4, 6
}

func sum(nums ...int) int {
	summa := 0

	for _, num := range nums {
		summa += num
	}

	return summa
}
func initSqn() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func factorial(n int) int {
	if n <= 1 {
		return n
	}

	return factorial(n-1) * n
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
