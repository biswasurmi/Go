package main

import "fmt"

func zero(xptr *int) {
	*xptr = 0
}
func one(xptr *int) {
	*xptr = 0
}
func square(x *int){
	*x = *x * *x
}
func swap(x *int, y *int){
	temp := *x
	*x = *y
	*y = temp
}
func main() {
	x := 5
	fmt.Println("Before function call x is ", x)
	zero(&x)
	fmt.Println("after function call x is ", x)
	/// using new for storing value in address
	y := new(int)
	one(y)
	fmt.Println(y) // prints address
	fmt.Println(*y) // prints value
	///example
	ex := 5
	square(&ex)
	fmt.Println(ex)
	xx := 5
	yy := 10
	swap(&xx, &yy)
	fmt.Println(xx, yy)
}
