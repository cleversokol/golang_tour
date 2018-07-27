package main

import "fmt"

func main() {


    i := 1 + 2
    i = i*5
    defer fmt.Println(i)
    fmt.Println("hello")
}
