package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 12
	m["k2"] = 43

	fmt.Println(m)
	delete(m, "k2")
	fmt.Println(m)
	clear(m)
	fmt.Println(m)

	_, k1 := m["k1"]

	fmt.Println(k1)

	m1 := map[string]int{"a": 1, "b": 3}

	m2 := map[string]int{"a": 1, "b": 3}

	fmt.Println(maps.Equal(m1, m2))
}
