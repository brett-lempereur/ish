package ish

import "image"

// Lookup tables using BT.601 constants for colour channel to grayscale
// conversion.
var redToGray = grayscaleTable(0.2989)
var greenToGray = grayscaleTable(0.5870)
var blueToGray = grayscaleTable(0.1140)

// Grayscale converts an image to the grayscale colour space, using optimised
// algorithms for common source colour spaces.
func Grayscale(img image.Image) *image.Gray {
	switch img := img.(type) {
	case *image.RGBA:
		return grayscaleRGBA(img)
	case *image.NRGBA:
		return grayscaleNRGBA(img)
	case *image.Gray:
		return img
	default:
		gray := image.NewGray(img.Bounds())
		for y := gray.Rect.Min.Y; y <= gray.Rect.Max.Y; y++ {
			for x := gray.Rect.Min.X; x <= gray.Rect.Max.X; x++ {
				gray.Set(x, y, img.At(x, y))
			}
		}
		return gray
	}
}

// grayscaleRGBA implements an optimised algorithm for converting images from
// the RGBA colour space to grayscale.
func grayscaleRGBA(img *image.RGBA) *image.Gray {
	gray := image.NewGray(img.Bounds())
	p := 0
	for q := 0; q < len(img.Pix); q += 4 {
		r := img.Pix[q]
		g := img.Pix[q+1]
		b := img.Pix[q+2]
		// Grayscale conversion using BT.601 lookup tables.
		gray.Pix[p] = uint8(redToGray[r] + greenToGray[g] + blueToGray[b])
		p++
	}
	return gray
}

// grayscaleNRGBA implements an optimised algorithm for converting images from
// the NRGBA colour space to grayscale.
func grayscaleNRGBA(img *image.NRGBA) *image.Gray {
	gray := image.NewGray(img.Bounds())
	p := 0
	for q := 0; q < len(img.Pix); q += 4 {
		// Alpha premultiplication.
		alpha := float32(img.Pix[q+3]) / 255.0
		r := uint8(float32(img.Pix[q]) * alpha)
		g := uint8(float32(img.Pix[q+1]) * alpha)
		b := uint8(float32(img.Pix[q+2]) * alpha)
		// Grayscale conversion using BT.601 lookup tables.
		gray.Pix[p] = uint8(redToGray[r] + greenToGray[g] + blueToGray[b])
		p++
	}
	return gray
}

// grayscaleTable implements an algorithm to generate lookup tables for colour
// to grayscale conversion.
func grayscaleTable(factor float64) [256]uint8 {
	var output [256]uint8
	for i := 0; i < 256; i++ {
		output[i] = uint8(float64(i) * factor)
	}
	return output
}
