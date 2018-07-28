package main

import "fmt"

func main() {

	input := 0
	fmt.Println("Enter how many fibonacci numbers do you want:")
	fmt.Scanln(&input)
	if input < 1 {
		return
	}

	fibs := fibonacci(input)
	for _, v := range fibs {
		fmt.Print(v, " ")
	}
	fmt.Println()
}

func fibonacci(length int) []int {

	result := []int{1, 1}

	for i := 2; i < length; i++ {
		next := result[i-2] + result[i-1]
		result = append(result, next)
	}
	return result
}
