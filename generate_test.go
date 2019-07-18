package palgen

import (
	"image/png"
	"os"
	"testing"
)

func TestGenerate(t *testing.T) {
	// Read a True Color PNG file
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
	// Read a True Color PNG file
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

	// Output a .gpl palette file
	err = Save(pal, "testdata/output.gpl", "From sample.png")
	if err != nil {
		t.Error(err)
	}
}

func TestSample2(t *testing.T) {
	// Read a True Color PNG file
	data, err := os.Open("testdata/splash.png")
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

	// Output a .gpl palette file
	err = Save(pal, "testdata/splash.gpl", "From splash.png")
	if err != nil {
		t.Error(err)
	}
}

func TestLarge(t *testing.T) {
	// Read a True Color PNG file
	data, err := os.Open("testdata/tm.png")
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

	// Output a .gpl palette file
	err = Save(pal, "testdata/tm.gpl", "From tm.png")
	if err != nil {
		t.Error(err)
	}
}
