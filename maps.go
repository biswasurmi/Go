package main

import "fmt"

func main() {
	// the below code will show error as the map is not initialized, we must need to initialized it first then assiging
	/*var x map[string]int
	x["key"] = 10
	fmt.Println(x["key"])*/
	x := make(map[string]int) // initialization
	x["key"] = 10             // assigning
	x["abc"] = 20
	x["lmn"] = 30
	//fmt.Print(x["key"])
	fmt.Println(x)
	delete(x, "key")     // deletion
	fmt.Println(x)       // full print of the map
	fmt.Println(x["ab"]) // return 0 for unassigned value
	name, ok := x["lmn"] // if value is present then ok is true else false
	fmt.Println(name, ok)
	if name, ok := x["ab"]; ok {
		fmt.Println(name, ok)
	}
	elements := map[string]string{
		"h":  "hydro",
		"he": "helium",
		"li": "lithium",
	}
	fmt.Println(elements)
	newEle := map[string]map[string]int{
		"h": map[string]int{
			"val":  10,
			"val1": 11,
		},
	}
	fmt.Println(newEle)
	if el, ok := newEle["h"]; ok {
		fmt.Println(el["val"], el["val1"])
	}
}
