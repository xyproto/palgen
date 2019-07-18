package palgen

import (
	"image"
	"image/png"
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	// Read a truecolor PNG file
	data, err := os.Open("testdata/tm.png")
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
	img = Convert(img, pal)

	// Convert the paletted image to an RGBA image
	imgRGBA := image.NewRGBA(img.Bounds())

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := img.At(x, y)
			imgRGBA.Set(x, y, c)
		}
	}

	// Output the truecolor image that represents the N-color image
	f, err := os.Create("image.png")
	if err != nil {
		t.Error(err)
	}

	if err := png.Encode(f, imgRGBA); err != nil {
		f.Close()
		t.Error(err)
	}

	if err := f.Close(); err != nil {
		t.Error(err)
	}

}
