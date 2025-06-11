package main
import "fmt"
func main(){
	var urmi [10] int // var array_name array_size data type
	urmi[5] = 10
	fmt.Println(urmi) // Print == Println

	for i := 1; i < 10; i++ {
		urmi[i] = i
	}
	var total float64 = 0
	for _, value := range urmi { // act like value : urmi , we can't use for i, value := range urmi as i is not used further in code and go doesn't allow unused var
		total += float64(value) // error if type conversion is not occurred
	}
	fmt.Println(total / float64(len(urmi)))

	x := [4]float64{
		1,
		2,
		3,
		4,  // this trailing ',' is necessary to remove elements easily
	}
	fmt.Println(x)
	var y []float64 // an array y is created with length of 0
	fmt.Print(y)
	fmt.Printf("\n")
	/*
	x := make([]float64, 5) create a slice of size 5 using make
	x := make([]float64, 5, 10) slice is of size 5 but the original array is of size 10
	x := []float64{1, 2, 3, 4, 5}
	x := arr[0:5] low is the index of whete to start the slice and high is the inces where to end it but doesn't include end index into the slice
	arr[0:5] = [1, 2, 3, 4, 5], arr[1:4] = [2, 3, 4] doesn't include indes 4 whose value is 5
	*/
	/*
	arr[0:] = arr[0: len(x)]
	arr[:5] = arr[0:5]
	arr[:] = arr[0:len(x)]
	*/
	//append
	/*slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)*/
	// copy
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)

	ex := make([]int, 3, 9)
	fmt.Println(ex)
	fmt.Println(len(ex))

	xx := [6]string{"a","b","c","d","e","f"}
	fmt.Println(xx[2:5])

	xxx := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}
	var choto int = 100;
	for i := 1; i < len(xxx); i++{
		if xxx[i] < choto {
			choto = xxx[i]
		}
	}
	fmt.Println(choto)
}