package datatypeV1

func Solution(S string) int {
	// Implement your solution here
	count := map[rune]int{
		'N': 0,
		'A': 0,
		'B': 0,
	}
	for _, c := range S {
		count[c]++
	}
	deletedCount := 0
	for {
		if count['A'] >= 3 && count['N'] >= 2 && count['B'] >= 1 {
			count['A'] -= 3
			count['N'] -= 2
			count['B'] -= 1
			deletedCount++
		} else {
			break
		}
	}
	return deletedCount
}

// func main() {

// 	println(Solution("NAABXXAN"))
// 	println(Solution("NAANAAXNABABYNNBZ"))

// }
