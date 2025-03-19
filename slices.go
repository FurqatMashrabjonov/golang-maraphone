package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("s: ", s, s == nil, len(s))

	s = make([]string, 3)
	fmt.Println("s: ", s, "len: ", len(s), "cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println(s, "len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f", "g")

	fmt.Println(s, "len: ", len(s))

	l := make([]string, len(s))
	copy(l, s)
	fmt.Println(l)
	fmt.Println(l[2:])
	fmt.Println(l[:1])
	fmt.Println(l[:])

	t := []string{"g", "h", "i"}
	fmt.Println(t)

	t2 := []string{"g", "h", "i"}

	if slices.Equal(t, t2) {
		fmt.Println("t and t2 are equal")
	}

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1

		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println(twoD)
}
