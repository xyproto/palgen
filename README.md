# palgen [![GoDoc](https://godoc.org/github.com/xyproto/palgen?status.svg)](http://godoc.org/github.com/xyproto/palgen) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/palgen)](https://goreportcard.com/report/github.com/xyproto/palgen)

Given an image, create a palette of N colors.

### Features and limitations

* Can generate palettes of N colors relatively quickly.
* The generated palette is not optimal, but it's okay.
* Can also export any given `color.Palette` to a GIMP `.gpl` palette file.
* The palette is generated by first grouping colors into N intensity levels and then use the median color of each group.
* Can convert truecolor `image.Image` images to indexed `image.Paletted` images.

### Example use

```go
// Read a PNG file
imageData, err := os.Open("input.png")
if err != nil {
	return err
}

// Decode the PNG image
img, err := png.Decode(imageData)
if err != nil {
	return err
}

// Generate a palette with 256 colors
pal, err := palgen.Generate(img, 256)
if err != nil {
	return err
}

// Output a .gpl palette file with the name "Untitled"
err = palgen.Save(pal, "output.gpl", "Untitled")
if err != nil {
	return err
}
```

### Included utilities

* `png256`, for converting a truecolor PNG image to an indexed PNG image, with a custom palette of 256 colors.
* `png2palette`, for extracting a palette from a truecolor PNG image and write the palette as an indexed 256 color PNG image.
* `png2gpl`, for extracting a palette from a truecolor PNG image and write the palette as a GIMP palette file (`.gpl`).

### Example palette extraction

| original | extracted 256 color palette |
| :---:    | :---:                       |
| ![png](testdata/splash.png) | ![png](testdata/splash_pal.png) |
| ![png](testdata/tm_small.png) | ![png](testdata/tm_small_pal.png) |

The palette can be extracted and saved as a PNG image, using `png2palette`, or as a GIMP palette, using `png2gpl`.

Palettes can be sorted by hue, luminance and chroma, using the HCL colorspace and the [go-colorful](https://github.com/lucasb-eyer/go-colorful) package, with the included `palgen.Sort` function. The above palettes are sorted with this method.

### Image Comparison

The palettes are generated by palgen

| original | 8 color palette | 16 color palette | 32 color palette | 64 color palette | 128 color palette | 256 color palette |
| :---: | :---: | :---: | :---: | :---: | :---: | :---: |
| ![png](testdata/splash.png)      | ![png](testdata/splash8.png)   | ![png](testdata/splash16.png)   | ![png](testdata/splash32.png)   | ![png](testdata/splash64.png)   | ![png](testdata/splash128.png)   | ![png](testdata/splash256.png)   |
| ![png](testdata/tm_small.png)      | ![png](testdata/tm_small8.png)   | ![png](testdata/tm_small16.png)   | ![png](testdata/tm_small32.png)   | ![png](testdata/tm_small64.png)   | ![png](testdata/tm_small128.png)   | ![png](testdata/tm_small256.png)   |

### General info

* Version: 3.0.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
