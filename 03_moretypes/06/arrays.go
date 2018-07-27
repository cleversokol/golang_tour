package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	var input int
	fmt.Scanln(&input)

	fmt.Println(input)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes)
	fmt.Println(isPrime(17))
}

func isPrime(input int) bool {
	for i := 2; i < input/2; i++ {
		if input%i == 0 {
			return false
		}
	}
	return true
}
