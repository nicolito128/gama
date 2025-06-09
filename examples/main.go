package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/nicolito128/gama"
)

var (
	length = flag.Int("length", 16, "get a palette with the given length")
)

func main() {
	flag.Parse()

	img, err := loadImage("./img.jpg")
	if err != nil {
		panic(err)
	}

	pl := gama.New(img)

	colors, err := pl.Quantify(*length)
	if err != nil {
		panic(err)
	}

	hexs := make([]string, len(colors))
	for i, c := range colors {
		hexs[i] = gama.ColorToHex(c, false)
	}

	fmt.Println("Palette: ", hexs)
}

func loadImage(filePath string) (image.Image, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	img, err := jpeg.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("could not decode image: %w", err)
	}

	return img, nil
}
