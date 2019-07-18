package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/palgen"
	"image"
	"image/color"
	"image/png"
	"os"
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

	// Sort the palette by hue, luminance and chroma
	palgen.Sort(pal)

	// Remove the alpha
	for i, c := range pal {
		rgba := color.RGBAModel.Convert(c).(color.RGBA)
		rgba.A = 255
		pal[i] = rgba
	}

	// The first color is now the darkest one
	darkIndex := uint8(0)

	// Let each row be blocks 32 wide
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
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
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
