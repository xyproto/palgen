package palgen

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	// Read a truecolor PNG file
	data, err := os.Open("testdata/splash.png")
	if err != nil {
		t.Error(err)
	}

	// Decode the PNG image
	img, err := png.Decode(data)
	if err != nil {
		t.Error(err)
	}

	N := 256

	// Generate a palette with N colors
	pal, err := Generate(img, N)
	if err != nil {
		t.Error(err)
	}

	// Convert the image to only use the given palette
	imgN := Convert(img, pal)

	_, ok := imgN.(image.PalettedImage)
	if !ok {
		t.Fatal("The image should be an image.PalettedImage")
	}

	// Output the indexed image
	f, err := os.Create("testdata/splash256.png")
	if err != nil {
		t.Error(err)
	}

	if err := png.Encode(f, imgN); err != nil {
		f.Close()
		t.Error(err)
	}

	if err := f.Close(); err != nil {
		t.Error(err)
	}

}
