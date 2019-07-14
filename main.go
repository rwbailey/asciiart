package main

import (
	"fmt"
	"image/jpeg"
	"math"
	"os"
	//"github.com/fatih/color"
)

func main() {
	pictureFile, err := os.Open("man.jpg")
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

	//pr := 0.21
	//pg := 0.72
	//pb := 0.07
	pr := 0.3
	pg := 0.59
	pb := 0.11

	var imageData [][]uint8

	for x := minX; x <= maxX; x++ {
		var vals []uint8
		for y := minY; y <= maxY; y++ {
			r, g, b, _ := picture.At(x,y).RGBA()
			rgb := uint8(pr*float64(r) + pg*float64(g) + pb*float64(b))
			//rgb := uint8((r + g + b)/3)
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
//color.Set(color.FgHiGreen)
	for y := minY; y <= maxY; y++ {
		for x := minX; x < maxX; x++ {
			//fmt.Println(y, x)
			fmt.Print(output[x][y])
			fmt.Print(output[x][y])
			fmt.Print(output[x][y])
		}
		fmt.Println("")
	}
//color.Unset()
}
