package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/palgen"
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
		fmt.Println("png256 1.0.0")
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

	// Convert the image to only use the given palette
	indexedImage, err := palgen.Convert(m)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// Output the indexed image
	if err := png.Encode(f, indexedImage); err != nil {
		f.Close()
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}
