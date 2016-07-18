package main

func main() {
	result, sum := add(1, 3, 4, 5, 6)
	println(result, sum)
}

func add(terms ...int) (numTerms int, sum int) {
	for _, term := range terms {
		sum += term
	}
	numTerms = len(terms)
	return
}
