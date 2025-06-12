package main

import ( // multiple import
	"fmt"
	"os"
)

func average(xs []float64) float64 {
	var tot float64 = 0.0
	for i := 0; i < len(xs); i++ {
		tot += xs[i]
	}
	return tot / float64(len(xs))
}
func f2() (r int) { // naming return type
	r = 1
	return
}
func multi() (int, int) { // returning multiple value
	return 5, 6
}
func add(nums ...int) int {
	var tot int = 0
	for _, value := range nums {
		tot += value
	}
	return tot
}
func main() {
	arr := []float64{1, 2, 3, 4, 5}
	fmt.Println(average(arr))
	fmt.Println(f2())
	x, y := multi()
	fmt.Println(x, y)
	/// variadic functions
	/* you can pass as many arguments as you wish, no need to fix the number and passes each element as a separate argument */
	fmt.Println(add(1, 2, 3))
	fmt.Println(add(3, 4))
	xs := []int{1, 2, 3, 4, 5}
	fmt.Println(add(xs...))

	/// closure => function within function
	xx := 0
	increment := func() int {
		xx++
		return xx
	}
	fmt.Println(increment())
	fmt.Println(increment())
	/// func in return type
	nextEven := evenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	/// recursion
	fmt.Println(factorial(5))
	/// defer => delays the execution of a function until the surrounding function exists and used for cleanup tasks (closing file, unlocking mutexes etc.) for multiple defer, follows LIPO
	defer second()
	first()
	third()
	defer first()
	third()

	f, err := os.Open("demo.go")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer f.Close()

	///panic and recover => PANIC stops the execution of current function when something goes very wrong, recover stops the panic but it works only when it is inside of any deferred function
	defer func() {
		str := recover()
		fmt.Println("Recovered from : ", str)
	}()
	panic("PANIC")

}
func first() {
	fmt.Println("first")
}
func second() {
	fmt.Println("second")
}
func third() {
	fmt.Println("third")
}
func evenGenerator() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}
func factorial(x int) int {
	if x == 1 {
		return 1
	} else {
		return x * factorial(x-1)
	}
}
