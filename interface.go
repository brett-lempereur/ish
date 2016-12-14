// Package ish implements a collection of perceptual hash algorithms for
// digital forensic image processing.
package ish

import "image"

// PerceptualHash is the interface implemented by perceptual hash algorithms.
type PerceptualHash interface {
	// Hash computes the perceptual hash of an image.
	Hash(img image.Image) ([]byte, error)
	// Length returns the length of hashes computed by this percepual hash
	// in bytes.
	Length() int
	// Distance calculates the hamming distance between two perceptual hashes.
	Distance(ha, hb []byte) int
}
