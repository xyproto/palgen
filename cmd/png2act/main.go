package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/palgen"
	"image/color"
	"image/png"
	"os"
)

const versionString = "png2act 1.0.1"

func main() {

	var (
		outputFilename string
		version        bool
		paletteName    string
	)

	flag.StringVar(&outputFilename, "o", "-", "output ACT palette")
	flag.BoolVar(&version, "v", false, "version")
	flag.StringVar(&paletteName, "n", "Untitled", "palette name")

	flag.Parse()

	if version {
		fmt.Println(versionString)
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println(versionString)
		fmt.Println("Extract an ACT palette from a given PNG file.")
		fmt.Println("Example use: png2act -o output.png input.png")
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

	// Output the ACT palette
	palgen.SaveACT(pal, "output.act")
}
