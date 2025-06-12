package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ { // must follow this structure with indentation and braces
		if i%2 == 0 {
			fmt.Println(i, "odd")
		} else {
			fmt.Println(i, "even")
		}
	}
	j := 1
	for j <= 5 {
		fmt.Print(j, " ") // fmt.Print(j, ' ') => integer is treated as unicode character rather than string thus shows unexpected value
		j += 1
	}
	fmt.Printf("\n") // fmt.Printf('\n') => error
	j = 1
	for j <= 5 {
		fmt.Printf("%d ", j)
		j += 1
	}
}
