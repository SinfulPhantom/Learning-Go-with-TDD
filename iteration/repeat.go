package iteration

func Repeat(character string, numberOfChars int) string {
	var repeated string

	for i := 0; i < numberOfChars; i++ {
		repeated += character
	}
	return repeated
}
