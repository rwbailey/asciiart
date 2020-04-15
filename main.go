package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"math"
	"os"
	//"github.com/fatih/color"
)

type picture struct {
	file image.Image
	minX int
	maxX int
	minY int
	maxY int
	pr   float32
	pg   float32
	pb   float32
}

func (p *picture) setDimensions() {
	p.minX = p.file.Bounds().Min.X
	p.maxX = p.file.Bounds().Max.X
	p.minY = p.file.Bounds().Min.Y
	p.maxY = p.file.Bounds().Max.Y
}

func (p *picture) setColourWeights() {
	p.pr = 0.3
	p.pg = 0.59
	p.pb = 0.11
}

func (p *picture) encode(file io.Reader) {
	var err error
	p.file, err = jpeg.Decode(file)
	if err != nil {
		fmt.Print("Error decoding file")
	}
}

func main() {
	var p picture

	pictureFile, err := os.Open("man.jpg")
	if err != nil {
		fmt.Print("Error loading file")
	}

	p.encode(pictureFile)
	p.setDimensions()
	p.setColourWeights()

	var imageData [][]uint8

	for x := p.minX; x <= p.maxX; x++ {
		var vals []uint8
		for y := p.minY; y <= p.maxY; y++ {
			r, g, b, _ := p.file.At(x, y).RGBA()
			rgb := uint8(p.pr*float32(r) + p.pg*float32(g) + p.pb*float32(b))
			vals = append(vals, rgb)
		}
		imageData = append(imageData, vals)
	}

	asciiString := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"

	var output [][]string

	if imageData != nil {
		for x := p.minX; x <= p.maxX; x++ {
			var pix []string
			for y := p.minY; y <= p.maxY; y++ {
				n := math.Ceil((float64(imageData[x][y]) / 255.0) * 64)
				//fmt.Print(n, ", ")
				pix = append(pix, string(asciiString[int(n)]))
			}
			output = append(output, pix)
		}
	}

	//color.Set(color.FgHiGreen)
	if output != nil {
		for y := p.minY; y <= p.maxY; y++ {
			for x := p.minX; x < p.maxX; x++ {
				//fmt.Println(y, x)
				fmt.Print(output[x][y])
				fmt.Print(output[x][y])
				fmt.Print(output[x][y])
			}
			fmt.Println("")
		}
	}
	//color.Unset()
}
