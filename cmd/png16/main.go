package main

import (
	"flag"
	"fmt"
	"github.com/xyproto/palgen"
	"image/png"
	"os"
)

const versionString = "png16 1.0.0"

func main() {

	var (
		outputFilename string
		version        bool
	)

	flag.StringVar(&outputFilename, "o", "-", "output PNG filename")
	flag.BoolVar(&version, "v", false, "version")

	flag.Parse()

	if version {
		fmt.Println(versionString)
		return
	}

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println(versionString)
		fmt.Println("Extract an indexed palette from a given PNG file and write a new PNG file which uses that palette.")
		fmt.Println("Example use: png16 -o output.png input.png")
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

	// Convert the image to only use the basic 16-color palette
	indexedImage, err := palgen.ConvertBasic(m)
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
