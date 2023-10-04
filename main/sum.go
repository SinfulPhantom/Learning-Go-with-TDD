package main

func Sum(numbers []int) int {
	result := 0

	for _, value := range numbers {
		result += value
	}
	return result
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

// Ended here https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#write-the-test-first-3
