package ish

import "testing"

// BenchmarkGrayscale benchmarks the grayscale conversion algorithm against
// different colourspaces and image sizes.
func BenchmarkGrayscale(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
	}{
		{"RGBA-4MP", "testdata/house-4MP-RGBA.png"},
		{"RGBA-8MP", "testdata/house-8MP-RGBA.png"},
		{"CMYK-4MP", "testdata/house-4MP-CMYK.png"},
		{"CMYK-8MP", "testdata/house-8MP-CMYK.png"},
		{"Gray-4MP", "testdata/house-4MP-Gray.png"},
		{"Gray-8MP", "testdata/house-8MP-Gray.png"},
	}
	for _, bm := range benchmarks {
		image, _, err := LoadFile(bm.filename)
		if err != nil {
			b.Error("Could not load", bm.filename, ":", err)
		}
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Grayscale(image)
			}
		})
	}
}
