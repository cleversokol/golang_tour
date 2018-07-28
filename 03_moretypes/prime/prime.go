package main

import "fmt"

func main() {

	input := 0
	primes := []int{2}

	fmt.Println("Enter how many primes do you want:")
	switch fmt.Scanln(&input); {
	case input < 1:
		return
	default:
		input--
		fmt.Print(primes[0], " ")
		defer fmt.Println()
	}

	for j := 3; input > 0; j += 2 {
		if isPrime(j) {
			fmt.Print(j, " ")
			primes = append(primes, j)
			input--
		}
	}
}

func isPrime(input int) bool {
	for i := 2; i < input/2; i++ {
		if input%i == 0 {
			return false
		}
	}
	return true
}
