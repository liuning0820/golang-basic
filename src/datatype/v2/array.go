package datatype

//切片
// Sum calculates the total from an slice of numbers.
func SumSlice(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}

	return sum
}
