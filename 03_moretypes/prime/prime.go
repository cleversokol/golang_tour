package main

import "fmt"

func main() {
	fmt.Println("Enter how many primes do you want:")
	var input int
	fmt.Scanln(&input)

	primes := []int{2}
	input--
	for j := 3; input > 0; j += 2 {
		if isPrime(j) {
			primes = append(primes, j)
			input--
		}
	}
	for i := range primes {
		fmt.Println(primes[i])
	}
	fmt.Println(primes)
}

func isPrime(input int) bool {
	for i := 2; i < input/2; i++ {
		if input%i == 0 {
			return false
		}
	}
	return true
}
