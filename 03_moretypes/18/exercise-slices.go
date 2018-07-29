package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	var s [][]uint8
	//var t []uint8

	for i := 0; i < dx; i++ {
		var t []uint8
		for j := 0; j < dy; j++ {
			t = append(t, 0) // "Black" square
		}
		s = append(s, t)
	}
	return s
}

func main() {
	pic.Show(Pic)
}
