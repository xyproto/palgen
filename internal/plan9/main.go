package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/palgen"
	"image/color/palette"
	"image/png"
	"os"
)

func main() {

	var (
		outputFilename string
		version        bool
		err            error
		notsorted      bool
		outputGPL      bool
	)

	flag.StringVar(&outputFilename, "o", "plan9.png", "output filename")
	flag.BoolVar(&version, "v", false, "version")
	flag.BoolVar(&notsorted, "u", true, "unsorted")
	flag.BoolVar(&outputGPL, "g", false, "output as a GPL palette")

	flag.Parse()

	if version {
		fmt.Println("plan9 1.0.0")
		os.Exit(0)
	}

	if outputGPL && outputFilename == "plan9.png" {
		outputFilename = "plan9.gpl"
	}

	// Get the Plan 9 palette
	pal := palette.Plan9

	if outputGPL {
		err := palgen.SaveGPL(pal, outputFilename, "plan9")
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
