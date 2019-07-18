package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/palgen"
	"image/color"
	"image/png"
	"os"
)

func main() {

	var (
		outputFilename string
		version        bool
		paletteName    string
	)

	flag.StringVar(&outputFilename, "o", "-", "output GPL palette")
	flag.BoolVar(&version, "v", false, "version")
	flag.StringVar(&paletteName, "n", "Untitled", "palette name")

	flag.Parse()

	if version {
		fmt.Println("png2gpl 1.0.0")
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

	// Output the GPL palette
	f.Write([]byte(palgen.GPL(pal, paletteName)))
	f.Close()
}
