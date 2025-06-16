package main

import (
	"fmt"
	"example.com/packages/math"
	m "example.com/packages/maxMin"
	
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)

	xs1 := []int{10, 20, 4, 1, 50}

	fmt.Println("Max is ", m.Max(xs1))

	fmt.Println("Min is ", m.Min(xs1))
}