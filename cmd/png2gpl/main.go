package main

import (
	"flag"
	"fmt"
	"image/color"
	"image/png"
	"os"

	"github.com/xyproto/palgen"
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
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// Decode the PNG file
	m, err := png.Decode(f)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	f.Close()

	// Prepare an output file
	var of *os.File

	// Prepare to output the new PNG data to either stdout or to file
	if outputFilename == "-" {
		of = os.Stdout
	} else {
		of, err = os.Create(outputFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
	}
	defer of.Close()

	// Create a new palette, with 256 colors (the rest of the code does not assume 256)
	pal, err := palgen.Generate(m, 256)
	if err != nil {
		of.Close()
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
	palString, err := palgen.GPL(pal, paletteName)
	if err != nil {
		of.Close()
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	of.Write([]byte(palString))
}
