package main

import (
	"fmt"
	"image/jpeg"
	"os"
	"math"
)

func main() {
	pictureFile, err := os.Open("Untitled.jpg")
	if err != nil {
		fmt.Print("Error loading file")
	}
	picture, err := jpeg.Decode(pictureFile)
	if err != nil {
		fmt.Print("Error decoding file")
	}

	minX := picture.Bounds().Min.X
	maxX := picture.Bounds().Max.X
	minY := picture.Bounds().Min.Y
	maxY := picture.Bounds().Max.Y

	var imageData [][]uint8

	for x := minX; x <= maxX; x++ {
		var vals []uint8
		for y := minY; y <= maxY; y++ {
			r, g, b, _ := picture.At(x,y).RGBA()
			rgb := uint8((int(r) + int(g) + int(b)) / 3)
			vals = append(vals, rgb)
		}
		imageData = append(imageData, vals)
	}

	asciiString := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

	//n := math.Ceil((255.0 / 255.0) * 65)
	//fmt.Println(n)
	//fmt.Println(string(asciiString[int(n - 1)]))

	var output [][]string

	for x := minX; x <= maxX; x++ {
		var pix []string
		for y := minY; y <= maxY; y++ {
			n := math.Ceil((float64(imageData[x][y]) / 255.0) * 64)
			//fmt.Print(n, ", ")
			pix = append(pix, string(asciiString[int(n)]))
		}
		output = append(output, pix)
	}

	for x := minX; x <= maxX; x++ {
		for y := minY; y <= maxY; y++ {
			fmt.Print(output[y][x])
			fmt.Print(output[y][x])
			fmt.Print(output[y][x])
		}
		fmt.Println("")
	}
}
