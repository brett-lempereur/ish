package ish

import (
	"image"

	"github.com/steakknife/hamming"
)

// AverageHash implements the average hash perceptual hashing algorithm
// described by Dr. Neal Krawetz.
type AverageHash struct {
	width  int
	height int
}

// NewAverageHash returns a new instance of AverageHash.
func NewAverageHash(width, height int) *AverageHash {
	if height <= 0 || width <= 0 {
		panic("Width and height must be positive")
	}
	if height&(height-1) != 0 || width&(width-1) != 0 {
		panic("Width and height must be powers of two")
	}
	return &AverageHash{width, height}
}

// Hash computes the difference hash of an image by shrinking it and comparing
// the relative brightness of pixels to the mean brightness of the image.
func (ah *AverageHash) Hash(img image.Image) ([]byte, error) {

	// Prepare the image by converting it to grayscale and resizing it, if
	// either or both of these operations have been performed, the cost of
	// these calls is negligible.
	ig := Grayscale(img)
	is, err := Resize(ig, ah.width, ah.height)
	if err != nil {
		return nil, err
	}

	// Compute the mean grayscale colour of the resized image.
	total := int(0)
	for i := 0; i < len(is.Pix); i++ {
		total += int(is.Pix[i])
	}
	mean := uint8(total / len(is.Pix))

	//
	hash := make([]byte, ah.Length())
	for i := uint(0); i < uint(len(is.Pix)); i++ {
		if is.Pix[i] >= mean {
			hash[i>>3] |= 1 << (i & 7)
		}
	}
	return hash, nil

}

// Length returns the length of the average hash in bytes.
func (ah *AverageHash) Length() int {
	return (ah.width * ah.height) / 8
}

// Distance returns the hamming distance between two difference hashes.
func (ah *AverageHash) Distance(ha, hb []byte) int {
	return hamming.Bytes(ha, hb)
}
