package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	o := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		o[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			o[y][x] = uint8((x + y) / 2)
		}
	}
	return o
}

func main() {
	pic.Show(Pic)
}
