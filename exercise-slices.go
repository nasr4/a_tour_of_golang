package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picSlice := make([][]uint8, dy)
	
	for i := 0; i < dy; i++ {
		picSlice[i] = make([]uint8, dx)
	}
	
	return picSlice
}

func main() {
	pic.Show(Pic)
}
