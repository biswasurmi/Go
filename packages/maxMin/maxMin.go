package math


func Max(xs []int) int {
	total := 0
	for i := 0; i < len(xs); i++ {
		if total < xs[i] {
			total = xs[i]
		}
	}
	return total
}

func Min(xs []int) int {
	total := 1000000000
	for i := 0; i < len(xs); i++ {
		if total > xs[i] {
			total = xs[i]
		}
	}
	return total
}