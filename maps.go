package main

import "fmt"

func main() {
	fmt.Println(getAverage([]float64{12, 21, 21, 31, 23, 123, 213, 123, 21, 321, 321, 321, 312}))
	fmt.Println(get())
}

func getAverage(arr []float64) (avg float64, sum float64) {
	sum = 0.0

	for _, value := range arr {
		sum += value
	}

	avg = sum / float64(len(arr))
	return
}

func get() (int, int) {
	return 3, 5
}
