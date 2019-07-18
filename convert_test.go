package palgen

import (
	"image"
	"image/png"
	"os"
	"strconv"
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

	// Convert the image to only use the given palette
	imgN, err := Convert(img)
	if err != nil {
		t.Error(err)
	}

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

func TestConvertFewColors(t *testing.T) {
	for _, N := range []int{8, 16, 32, 64, 128} {

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

		// Generate a palette with N colors
		pal, err := Generate(img, N)
		if err != nil {
			t.Error(err)
		}

		if len(pal) != N {
			t.Fatalf("The palette should be %d long, but it is %d long.\n", N, len(pal))
		}

		// Convert the image to only use the given palette
		imgN, err := ConvertCustom(img, pal)
		if err != nil {
			t.Error(err)
		}

		_, ok := imgN.(image.PalettedImage)
		if !ok {
			t.Fatal("The image should be an image.PalettedImage")
		}

		// Output the indexed image
		f, err := os.Create("testdata/splash" + strconv.Itoa(N) + ".png")
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
}

func TestConvertGeneral(t *testing.T) {
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

	// Convert the image to only use the given palette
	imgN, err := ConvertGeneral(img)
	if err != nil {
		t.Error(err)
	}

	_, ok := imgN.(image.PalettedImage)
	if !ok {
		t.Fatal("The image should be an image.PalettedImage")
	}

	// Output the indexed image
	f, err := os.Create("testdata/splash_standard256.png")
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

func TestConvertMountain(t *testing.T) {
	for _, N := range []int{8, 16, 32, 64, 128, 256} {

		// Read a truecolor PNG file
		data, err := os.Open("testdata/tm_small.png")
		if err != nil {
			t.Error(err)
		}

		// Decode the PNG image
		img, err := png.Decode(data)
		if err != nil {
			t.Error(err)
		}

		// Generate a palette with N colors
		pal, err := Generate(img, N)
		if err != nil {
			t.Error(err)
		}

		if len(pal) != N {
			t.Fatalf("The palette should be %d long, but it is %d long.\n", N, len(pal))
		}

		// Convert the image to only use the given palette
		imgN, err := ConvertCustom(img, pal)
		if err != nil {
			t.Error(err)
		}

		_, ok := imgN.(image.PalettedImage)
		if !ok {
			t.Fatal("The image should be an image.PalettedImage")
		}

		// Output the indexed image
		f, err := os.Create("testdata/tm_small" + strconv.Itoa(N) + ".png")
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
}
