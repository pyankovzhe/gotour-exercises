package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	points := make([]uint8, dx)
	for i := range picture {
		for j := range points {
			points[j] = uint8((i + j) / 2)
		}
		picture[i] = points
	}

	return picture
}

func main() {
	pic.Show(Pic)
}
