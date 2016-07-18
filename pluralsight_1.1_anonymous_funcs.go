package main

func main() {
	addFunc := func(terms ...int) (numTerms int, sum int) {
		for _, term := range terms {
			sum += term
		}
		numTerms = len(terms)
		return
	}
	x, y := addFunc(1, 2, 3, 4, 5 ,6)
	println(x, y)
}
