# Gama

A Go library to obtain the color palette of an image.

## Requirements

* `Go v1.24+`

## Getting started

    go get -u github.com/nicolito128/gama

## Usage

```go
package main

import (
	"fmt"
	"image/png"
	"os"

	"github.com/nicolito128/gama"
)

const length = 16

func main() {
    file, err := os.Open("/path/to/example.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
	    panic(err)
    }

    pl := gama.New(img)

	colors, err := pl.Quantify(length)
	if err != nil {
		panic(err)
	}

	hexs := make([]string, len(colors))
	for i, c := range colors {
		hexs[i] = gama.ColorToHex(c, false)
	}

	fmt.Println("Palette: ", hexs)

}
```

You can also consult the folder [examples/](./examples/).

