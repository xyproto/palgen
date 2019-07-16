package palgen

import (
	"image/png"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	// Read a truecolor PNG file
	data, err := os.Open("testdata/sample.png")
	if err != nil {
		t.Error(err)
	}

	// Decode the PNG image
	img, err := png.Decode(data)
	if err != nil {
		t.Error(err)
	}

	// Generate a palette with 2 colors
	_, err = Generate(img, 2)
	if err != nil {
		t.Error(err)
	}
	//fmt.Println("Palette with 2 colors", pal)

	// Generate a palette with 256 colors
	pal, err := Generate(img, 256)
	if err != nil {
		t.Error(err)
	}
	//fmt.Println("Palette with 256 colors", pal)

	// Output a .gpl palette file in testdata/output.gpl
	err = Write(pal, "testdata/output.gpl", "Testing123")
	if err != nil {
		t.Error(err)
	}
}
