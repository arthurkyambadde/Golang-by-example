package Sums

func Sum(numbers ...int) int {
	var total int = 0

	for _, num := range numbers {
		total += num
	}
	return total
}
