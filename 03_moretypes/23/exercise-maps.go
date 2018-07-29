package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	t := strings.Fields(s)
	answer := make(map[string]int)

	for i, word := range t {
		count, ok := answer[word]
		if ok {
			answer[word] = count + 1
		} else {
			answer[word] = 1
		}
		fmt.Println(i, word)
	}
	//return map[string]int{"x": 1}
	return answer
}

func main() {
	wc.Test(WordCount)
}
