package ish

import (
	"image"
	"testing"
)

// BenchmarkResize benchmarks the image resizing algorithm against different
// source and target image dimensions.
func BenchmarkResize(b *testing.B) {
	benchmarks := []struct {
		name          string
		filename      string
		width, height int
	}{
		{"8MP-128x128", "testdata/house-8MP-Gray.png", 128, 128},
		{"8MP-64x64", "testdata/house-8MP-Gray.png", 64, 64},
		{"8MP-32x32", "testdata/house-8MP-Gray.png", 32, 32},
		{"8MP-16x16", "testdata/house-8MP-Gray.png", 16, 16},
		{"8MP-8x8", "testdata/house-8MP-Gray.png", 8, 8},
		{"4MP-128x128", "testdata/house-4MP-Gray.png", 128, 128},
		{"4MP-64x64", "testdata/house-4MP-Gray.png", 64, 64},
		{"4MP-32x32", "testdata/house-4MP-Gray.png", 32, 32},
		{"4MP-16x16", "testdata/house-4MP-Gray.png", 16, 16},
		{"4MP-8x8", "testdata/house-4MP-Gray.png", 8, 8},
		{"32x32-32x32", "testdata/house-32x32-Gray.png", 32, 32},
	}
	for _, bm := range benchmarks {
		img, _, err := LoadFile(bm.filename)
		if err != nil {
			b.Error("Could not load", bm.filename, ":", err)
		}
		gray, ok := img.(*image.Gray)
		if !ok {
			b.Error("Image", bm.filename, "is not grayscale")
		}
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := Resize(gray, bm.width, bm.height)
				if err != nil {
					b.Error("Could not resize", bm.filename, ":", err)
				}
			}
		})
	}
}
