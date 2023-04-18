package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
	"strconv"

	"github.com/xyproto/palgen"
)

const versionString = "pngn 1.0.1"

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
		fmt.Println("Also takes a number for how many colors to reduce the palette to. 256 is the default.")
		fmt.Println("Example use: pngn -o output.png input.png 128")
		os.Exit(1)
	}

	inputFilename := args[0]

	// The number of colors to reduce the image to
	n := 256
	if len(args) > 2 {
		if x, err := strconv.Atoi(args[1]); err == nil { // success
			n = x
		}
	}

	// Open the PNG file
	f, err := os.Open(inputFilename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}

	// Decode the PNG file
	img, err := png.Decode(f)
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

	// Reduce the image to only use a general N color palette
	indexedImage, err := palgen.Reduce(img, n)
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
