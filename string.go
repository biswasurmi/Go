package main

import "fmt"

func main(){
	fmt.Println(len("Urmi")) // print with a newline
	fmt.Println("hello world"[1]) // print asci value of string index [1]
	fmt.Printf("%c\n", "hello world"[1]) // print the character itself of the string
	fmt.Println("hello " + "world") // concatenate
}