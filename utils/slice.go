package utils

func SumInts(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func FindMinInt(fn func(i int) int, rng []int) int {
	min := fn(rng[0])
	for i := 1; i < len(rng); i++ {
		val := fn(rng[i])
		if val < min {
			min = val
		}
	}
	return min
}

func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
