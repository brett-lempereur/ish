package ish

import (
	"image"

	"github.com/steakknife/hamming"
)

// DifferenceHash implements the difference hash perceptual hashing algorithm
// described by Dr. Neal Krawetz.
type DifferenceHash struct {
	width  int
	height int
}

// NewDifferenceHash returns a new instance of DifferenceHash.
func NewDifferenceHash(width, height int) *DifferenceHash {
	if height <= 0 || width <= 0 {
		panic("Width and height must be positive")
	}
	if height&(height-1) != 0 || width&(width-1) != 0 {
		panic("Width and height must be powers of two")
	}
	return &DifferenceHash{width, height}
}

// Hash computes the difference hash of an image by shrinking it and comparing
// the relative brightness of pixels.
func (dh *DifferenceHash) Hash(img image.Image) ([]byte, error) {

	// Prepare the image by converting it to grayscale and resizing it, if
	// either or both of these operations have been performed, the cost of
	// these calls is negligible.
	ig := Grayscale(img)
	is, err := Resize(ig, dh.width+1, dh.height)
	if err != nil {
		return nil, err
	}

	// Compute the difference hash by scanning the image horizontally and
	// comparing the brightness of pixels.
	hash := make([]byte, dh.Length())
	for y := 0; y < dh.height; y++ {
		for x := 0; x < dh.width; x++ {
			i := uint(y*dh.width + x)
			if is.Pix[i] < is.Pix[i+1] {
				hash[i>>3] |= 1 << (i & 7)
			}
		}
	}
	return hash, nil

}

// Length returns the length of the difference hash in bytes.
func (dh *DifferenceHash) Length() int {
	return (dh.width * dh.height) / 8
}

// Distance returns the hamming distance between two difference hashes.
func (dh *DifferenceHash) Distance(ha, hb []byte) int {
	return hamming.Bytes(ha, hb)
}
