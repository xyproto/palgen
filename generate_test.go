package palgen

import (
	"fmt"
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
	err = SaveGPL(pal, "testdata/output.gpl", "From sample.png")
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
	err = SaveGPL(pal, "testdata/splash.gpl", "From splash.png")
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
	err = SaveGPL(pal, "testdata/tm.gpl", "From tm.png")
	if err != nil {
		t.Error(err)
	}
}

func TestSmallImagePalette(t *testing.T) {
	// Read images and try to generate palettes of exactly 2 and 4 colors
	for _, n := range []int{2, 4} {
		for _, imageName := range []string{"rainforest", "splash", "tm"} {
			data, err := os.Open(fmt.Sprintf("testdata/%s.png", imageName))
			if err != nil {
				t.Errorf("Failed to open image file: %v", err)
			}
			defer data.Close()

			img, err := png.Decode(data)
			if err != nil {
				t.Errorf("Failed to decode PNG image: %v", err)
			}

			// Generate a palette of N colors
			pal, err := Generate(img, n)
			if err != nil {
				t.Errorf("Failed to generate a palette of %d colors: %v", n, err)
			}

			// Check if the generated palette has exactly N colors
			if len(pal) != n {
				t.Errorf("Generated palette for testdata/%s.png has %d colors, expected %d", imageName, len(pal), n)
			}
		}
	}
}
