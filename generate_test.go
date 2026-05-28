package palgen

import (
	"fmt"
	"image"
	"image/color"
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

func TestGenerateOneColor(t *testing.T) {
	// Read a True Color PNG file
	data, err := os.Open("testdata/sample.png")
	if err != nil {
		t.Fatal(err)
	}
	defer data.Close()

	// Decode the PNG image
	img, err := png.Decode(data)
	if err != nil {
		t.Fatal(err)
	}

	// Generate a palette with 1 color
	pal, err := Generate(img, 1)
	if err != nil {
		t.Fatalf("Generate(img, 1) returned error: %v", err)
	}
	if len(pal) < 1 {
		t.Error("Generate(img, 1) returned an empty palette")
	}
}

func TestMedianNoOverflow(t *testing.T) {
	// Two bright colors where the sum of each component exceeds 255
	c1 := color.RGBA{250, 240, 230, 255}
	c2 := color.RGBA{240, 230, 220, 255}
	colors := []color.Color{c1, c2}

	median, err := Median(colors)
	if err != nil {
		t.Fatal(err)
	}

	rgba := median.(color.RGBA)
	wantR, wantG, wantB, wantA := uint8(245), uint8(235), uint8(225), uint8(255)
	if rgba.R != wantR || rgba.G != wantG || rgba.B != wantB || rgba.A != wantA {
		t.Errorf("Median overflow: got (%d,%d,%d,%d), want (%d,%d,%d,%d)",
			rgba.R, rgba.G, rgba.B, rgba.A, wantR, wantG, wantB, wantA)
	}
}

func TestGenerateSmallNBrightColors(t *testing.T) {
	// Create a 4x4 image with two bright colors that would trigger overflow
	img := image.NewRGBA(image.Rect(0, 0, 4, 4))
	bright1 := color.RGBA{220, 200, 180, 255}
	bright2 := color.RGBA{200, 180, 220, 255}
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if x < 2 {
				img.Set(x, y, bright1)
			} else {
				img.Set(x, y, bright2)
			}
		}
	}

	// Generate a palette with 2 colors
	pal, err := Generate(img, 2)
	if err != nil {
		t.Fatalf("Generate(img, 2) returned error: %v", err)
	}
	if len(pal) == 0 {
		t.Fatal("Generate(img, 2) returned an empty palette")
	}

	// Verify no color component has wrapped around due to overflow.
	// Both input colors have all components >= 180, so no palette color
	// should have a component below 90 (the midpoint of the dimmest pair).
	for i, c := range pal {
		rgba := color.RGBAModel.Convert(c).(color.RGBA)
		if rgba.R < 90 || rgba.G < 90 || rgba.B < 90 {
			t.Errorf("palette[%d] = (%d,%d,%d): likely uint8 overflow in median averaging",
				i, rgba.R, rgba.G, rgba.B)
		}
	}
}

func TestColorLengthOrdering(t *testing.T) {
	// A brighter color must have a greater length than a dimmer one
	bright := color.RGBA{255, 255, 255, 255}
	dim := color.RGBA{10, 10, 10, 255}

	brightLen := colorLength(bright)
	dimLen := colorLength(dim)
	if brightLen <= dimLen {
		t.Errorf("colorLength(255,255,255)=%f <= colorLength(10,10,10)=%f", brightLen, dimLen)
	}
}

func TestDeduplicateDarkColors(t *testing.T) {
	// Use a real image to test deduplication with a meaningful palette size
	data, err := os.Open("testdata/splash.png")
	if err != nil {
		t.Fatal(err)
	}
	defer data.Close()

	img, err := png.Decode(data)
	if err != nil {
		t.Fatal(err)
	}

	pal, err := Generate(img, 256)
	if err != nil {
		t.Fatalf("Generate returned error: %v", err)
	}

	// Count how many colors in the palette are very dark (gray < 13, i.e. 5% of 255)
	darkCount := 0
	for _, c := range pal {
		gc := color.GrayModel.Convert(c).(color.Gray)
		if gc.Y < 13 {
			darkCount++
		}
	}

	// The 10% darkest in a 256-color palette is ~25 entries.
	// After deduplication, near-identical darks within 5% lightness should be merged,
	// so we should not see an excessive number of very dark entries.
	t.Logf("palette has %d very dark colors (gray < 13) out of %d total", darkCount, len(pal))
	if darkCount > len(pal)/10 {
		t.Errorf("palette has %d very dark colors (gray < 13), expected at most %d (10%% of palette)", darkCount, len(pal)/10)
	}
}
