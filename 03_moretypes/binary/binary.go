package main

import "fmt"

func main() {
	for i := uint(0); i < 10; i++ {
		fmt.Println(2 << i)
		fmt.Println(3 << i)
	}

	fmt.Println(250 << uint(2))   // 250 * 4
	fmt.Println(16384 >> uint(4)) // 16384 / 16
}
