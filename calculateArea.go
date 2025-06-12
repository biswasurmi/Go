package main
import (
	"fmt"
	"math"
)
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a * a + b * b)
}
func rectangle(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x2, y2)
	w := distance(x1, y1, x2, y2)
	return l * w
}
func circle(r float64) float64 {
	return math.Pi * r * r
}
func main (){
	var x1, x2, y1, y2, r float64 = 10, 5, 10, 5, 10
	fmt.Println(rectangle(x1, y1, x2, y2), circle(r))
}
