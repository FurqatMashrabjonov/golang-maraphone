package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("empty:", a)

	a[4] = 100
	fmt.Println("get:", a)
	fmt.Println("set:", a[4])

	fmt.Println("len:", len(a))

	b := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(b)
	//f you specify the index with :, the elements in between will be zeroed.
	b = [...]int{100, 3: 400, 500, 600}
	fmt.Println(b)

	var twoD [2][3]int

	for i := 0; i < 2; i++ {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println(twoD)
}
