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

	// Generate a palette with 4 colors
	_, err = Generate(img, 4)
	if err != nil {
		t.Error(err)
	}
	//fmt.Println("Palette with 2 colors", pal)
}

func TestSample(t *testing.T) {
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

	// Generate a palette with 256 colors
	pal, err := Generate(img, 256)
	if err != nil {
		t.Error(err)
	}

	// Output a .gpl palette file in testdata/output1.gpl
	err = Save(pal, "testdata/output.gpl", "From sample.png")
	if err != nil {
		t.Error(err)
	}
}

func TestSample2(t *testing.T) {
	// Read a truecolor PNG file
	data, err := os.Open("testdata/sample2.png")
	if err != nil {
		t.Error(err)
	}

	// Decode the PNG image
	img, err := png.Decode(data)
	if err != nil {
		t.Error(err)
	}

	// Generate a palette with 256 colors
	pal, err := Generate(img, 256)
	if err != nil {
		t.Error(err)
	}

	// Output a .gpl palette file in testdata/output2.gpl
	err = Save(pal, "testdata/output2.gpl", "From sample2.png")
	if err != nil {
		t.Error(err)
	}
}
