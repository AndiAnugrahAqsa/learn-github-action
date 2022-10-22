package main

func Add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

func Subtract(numbers ...int) int {
	result := numbers[0]
	numbers = append([]int{}, numbers[1:]...)
	for _, number := range numbers {
		result -= number
	}

	return result
}

func main() {
	Add(3, 5, 1)
	Subtract(4, 3, 2)
}
