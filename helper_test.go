package ish

import (
	"bytes"
	"io/ioutil"
	"testing"
)

// BenchmarkLoadFile benchmarks the load file routine against a collection
// of image sizes and file formats.
func BenchmarkLoadFile(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
	}{
		{"RGBA-4MP-PNG", "testdata/house-4MP-RGBA.png"},
		{"RGBA-8MP-PNG", "testdata/house-8MP-RGBA.png"},
		{"RGBA-4MP-JPEG", "testdata/house-4MP-RGBA.jpeg"},
		{"RGBA-8MP-JPEG", "testdata/house-8MP-RGBA.jpeg"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, _, err := LoadFile(bm.filename)
				if err != nil {
					b.Error("Could not load", bm.filename, ":", err)
				}
			}
		})
	}
}

// BenchmarkDecodeFile benchmarks the decoding routine against a collection
// of image sizes and file formats.
func BenchmarkDecodeFile(b *testing.B) {
	benchmarks := []struct {
		name     string
		filename string
	}{
		{"RGBA-4MP-PNG", "testdata/house-4MP-RGBA.png"},
		{"RGBA-8MP-PNG", "testdata/house-8MP-RGBA.png"},
		{"RGBA-4MP-JPEG", "testdata/house-4MP-RGBA.jpeg"},
		{"RGBA-8MP-JPEG", "testdata/house-8MP-RGBA.jpeg"},
	}
	for _, bm := range benchmarks {
		data, err := ioutil.ReadFile(bm.filename)
		if err != nil {
			b.Error("Could not load", bm.filename, ":", err)
		}
		reader := bytes.NewReader(data)
		b.Run(bm.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				_, _, err := Decode(reader)
				if err != nil {
					b.Error("Could not decode", bm.filename, ":", err)
				}
				reader.Reset(data)
			}
		})
	}
}
