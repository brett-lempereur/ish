package ish

import "image"

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
		// TODO: Benchmarks show this is surprisingly fast for CMYK images on
		// Go1.7/amd64, maybe it's the SSA-backend?  Either way, this might be
		// the fastest path for all images now.  Poke around in the libraries
		// and check the dissassembly of these functions.
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
		r := float32(img.Pix[q])
		g := float32(img.Pix[q+1])
		b := float32(img.Pix[q+2])
		// Grayscale conversion using BT.601 constants.
		gray.Pix[p] = uint8(0.2989*r + 0.5870*g + 0.1140*b)
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
		r := float32(img.Pix[q]) * alpha
		g := float32(img.Pix[q+1]) * alpha
		b := float32(img.Pix[q+2]) * alpha
		// Grayscale conversion using BT.601 constants.
		gray.Pix[p] = uint8(0.2989*r + 0.5870*g + 0.1140*b)
		p++
	}
	return gray
}
