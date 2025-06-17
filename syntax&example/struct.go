package main

import (
	"fmt"
	"math"
)

// / circle
type circle struct {
	x, y, r float64
}

func circleArea(c circle) float64 {
	return math.Pi * c.r * c.r
}
func circleAreaWithPointer(c *circle) float64 {
	return math.Pi * c.r * c.r
}

// / rectangle
type rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}
func (r *rectangle) area() float64 { // func (struct_name) func_name func_return_type
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func main() {
	var c circle
	// c := new(circle)
	//c := circle{x : 0, y: 0, r: 5} // this or
	c = circle{0, 0, 5} // this
	fmt.Println(c.x, c.y, c.r)
	c.x = 10
	c.y = 5
	c.r = 5
	fmt.Println(circleArea(c))
	fmt.Println(circleAreaWithPointer(&c))
	r := rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}
