package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/palgen"
	"image"
	"image/color"
	"image/png"
	"os"
	"sort"
)

func main() {

	var (
		outputFilename string
		version        bool
	)

	flag.StringVar(&outputFilename, "o", "-", "output PNG filename")
	flag.BoolVar(&version, "v", false, "version")

	flag.Parse()

	if version {
		fmt.Println("png2palette 1.0.0")
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "An input PNG filename is required.\n")
		os.Exit(1)
	}

	inputFilename := args[0]

	// Open the PNG file
	f, err := os.Open(inputFilename)
	m, err := png.Decode(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	f.Close()

	// Prepare to output the new PNG data to either stdout or to file
	if outputFilename == "-" {
		f = os.Stdout
	} else {
		f, err = os.Create(outputFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	// Create a new palette, with 256 colors (the rest of the code does not assume 256)
	pal, err := palgen.Generate(m, 256)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// Remove the alpha
	for i, c := range pal {
		rgba := color.RGBAModel.Convert(c).(color.RGBA)
		rgba.A = 255
		pal[i] = rgba
	}

	// Sort the palette
	spal := palgen.SortablePalette(pal)
	sort.Sort(spal)
	pal = color.Palette(spal)

	// Find the darkest color
	darkIndex := uint8(0)
	s := 0
	minsum := 255 + 255 + 255
	for i, c := range pal {
		rgba := color.RGBAModel.Convert(c).(color.RGBA)
		s = int(rgba.R + rgba.G + rgba.B)
		if s < minsum {
			minsum = s
			darkIndex = uint8(i)
		}
	}

	// Each row should be 32 wide
	w := 32
	h := len(pal) / w
	leftover := len(pal) % w
	if leftover > 0 {
		h++
	}

	// "pixel" width and height
	pw := 16
	ph := 16

	upLeft := image.Point{0, 0}
	lowRight := image.Point{w * pw, h * ph}

	// Create a new image, where a square is painted per color in the palette
	palImage := image.NewPaletted(image.Rectangle{upLeft, lowRight}, pal)

	// Set color for each pixel.
	colorIndex := uint8(0)
OUT:
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			for by := 0; by <= ph; by++ {
				for bx := 0; bx <= pw; bx++ {
					if bx == 0 || by == 0 || bx == pw-1 || by == ph-1 {
						palImage.SetColorIndex((x*pw)+bx, (y*ph)+by, darkIndex)
					} else {
						palImage.SetColorIndex((x*pw)+bx, (y*ph)+by, colorIndex)
					}
				}
			}
			colorIndex++
			if int(colorIndex) >= len(pal) {
				break OUT
			}
		}
	}

	// Output the generated image
	if err := png.Encode(f, palImage); err != nil {
		f.Close()
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
