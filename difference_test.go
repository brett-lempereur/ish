package ish

import "testing"

// BenchmarkDifferenceHash benchmarks the difference hash algorithm against
// different image sizes and hash lengths.
func BenchmarkDifferenceHash(b *testing.B) {
	benchmarks := []struct {
		name          string
		filename      string
		width, height int
	}{
		{"8MP-32x32", "testdata/house-8MP-RGBA.png", 32, 32},
		{"8MP-16x16", "testdata/house-8MP-RGBA.png", 16, 16},
		{"8MP-8x8", "testdata/house-8MP-RGBA.png", 8, 8},
		{"4MP-32x32", "testdata/house-4MP-RGBA.png", 32, 32},
		{"4MP-16x16", "testdata/house-4MP-RGBA.png", 16, 16},
		{"4MP-8x8", "testdata/house-4MP-RGBA.png", 8, 8},
	}
	for _, bm := range benchmarks {
		img, _, err := LoadFile(bm.filename)
		if err != nil {
			b.Error("Could not load", bm.filename, ":", err)
		}
		b.Run(bm.name, func(b *testing.B) {
			dh := NewDifferenceHash(bm.width, bm.height)
			for i := 0; i < b.N; i++ {
				_, err := dh.Hash(img)
				if err != nil {
					b.Error("Could not hash", bm.filename, ":", err)
				}
			}
		})
	}
}
