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
	)

	flag.StringVar(&outputFilename, "o", "websafe.png", "output PNG filename")
	flag.BoolVar(&version, "v", false, "version")
	flag.BoolVar(&notsorted, "u", true, "unsorted")

	flag.Parse()

	if version {
		fmt.Println("websafe 1.0.1")
		os.Exit(0)
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

	// Get the WebSafe palette
	pal := palette.WebSafe

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
