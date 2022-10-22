package main

func Add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

func Subtract(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result -= number
	}

	return result
}

func main() {
	Add(3, 5, 1)
	Subtract(4, 3, 2)
}
