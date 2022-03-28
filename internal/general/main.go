package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"

	"github.com/xyproto/palgen"
)

func main() {

	var (
		outputFilename string
		version        bool
		err            error
		notsorted      bool
		outputGPL      bool
	)

	flag.StringVar(&outputFilename, "o", "general.png", "output filename")
	flag.BoolVar(&version, "v", false, "version")
	flag.BoolVar(&notsorted, "u", true, "unsorted")
	flag.BoolVar(&outputGPL, "g", false, "output as a GPL palette")

	flag.Parse()

	if version {
		fmt.Println("general 1.0.1")
		os.Exit(0)
	}

	if outputGPL && outputFilename == "general.png" {
		outputFilename = "general.gpl"
	}

	// Get the general palette
	pal := palgen.GeneralPalette()

	if outputGPL {
		err := palgen.SaveGPL(pal, outputFilename, "general")
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		return // success
	}

	// Prepare to output the new PNG data to either stdout or to file
	f := os.Stdout
	if outputFilename != "-" {
		f, err = os.Create(outputFilename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	// Sort the palette by hue, luminance and chroma
	if !notsorted {
		palgen.Sort(pal)
	}

	// Render the palette as an image
	palImage := palgen.Render(pal)
	// Output the rendered image
	if err := png.Encode(f, palImage); err != nil {
		f.Close()
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
