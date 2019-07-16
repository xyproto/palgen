# palgen [![GoDoc](https://godoc.org/github.com/xyproto/palgen?status.svg)](http://godoc.org/github.com/xyproto/palgen) [![Go Report Card](https://goreportcard.com/badge/github.com/xyproto/palgen)](https://goreportcard.com/report/github.com/xyproto/palgen)

Given an image, create a palette of N colors.

## Features and limitations

* Can generate palettes of N colors relatively quickly.
* The generated palette is not optimal, but it's okay.
* Can also export any given `color.Palette` to a GIMP `.gpl` palette file.

## Image Comparison

| 64x64 image, 8-bit/color RGB  | 64x64 image, 256 color palette selected by `palgen` |
| --------------------------------- | -------------------------------- |
| ![png](testdata/sample2.png)      | ![png](testdata/sample2_converted.png)   |

## General info

* Version: 2.0.0
* License: MIT
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
