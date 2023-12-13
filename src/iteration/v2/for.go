package iteration

func Repeat(character string, count int) string {
	var repeated string
	var repeatCount = count
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
