package main

import "fmt"

// globalVariable := 10 // will show error as := this is only for local function
var globalVariable = 10 // you must use var for global variable
func main() {
	var x string // var variable_name variable_type
	x = "Hlw world!"
	fmt.Println(x)
	var y string = "hi world!@"
	fmt.Println(y)
	var z string = x + y
	fmt.Println(z)
	fmt.Println(x == y)
	p := 5 // : is for assigning value and no need to declare data type
	fmt.Println(p)
	var q = "Hlw world" // you can't use := here as you have declared 'var' here but no need to declare data type
	fmt.Println(q)
	myName := "Urmi"
	fmt.Println("My name is ", myName)
	fmt.Println(globalVariable)
	f() // calling another function
	const con = 10
	// con = 12 // error as assigning is not allowed for constants

	var (
		a        = 5
		b string = "hlw"
	)
	fmt.Println(a, b)
	//example()
	celToFar()
}
func f() {
	fmt.Println(globalVariable)
}
func example() {
	var input float64
	fmt.Print("enter a float number ")
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
	output += 1
	fmt.Println(output)
}
func celToFar() {
	var far float64
	fmt.Scanf("%f", &far)
	var cel float64 = ((far - 32) * (5.0 / 9))
	fmt.Println(cel)

}
