package main
import ("fmt")

func digits (number int, digitch chan int) {
	for number != 0 {
		digitch <- number %10
		number /= 10
	}
	close(digitch)
}

func calcSquare (number int, squarech chan int) {
	sum := 0
	dch := make(chan int)
	go digits (number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squarech <- sum
}

func calccube (number int, cubech chan int) {
	sum := 0
	dch := make(chan int)
	go digits (number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubech <- sum
}

func main() {
	number := 589
	squarech := make(chan int)
	cubech := make(chan int)
	go calcSquare(number, squarech)
	go calccube (number, cubech)
	fmt.Println("Final Output :", <- squarech + <- cubech)
}