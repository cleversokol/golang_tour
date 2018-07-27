package main

import "fmt"

type Vertex struct {
    X int
    Y int
    introduction string
}

func main() {
    v := Vertex{1, 2, "Hello, World!\n"}
    v.X = 4
    fmt.Println(v)
}