package ish

import (
	"image"
	"io"
	"os"

	// Register all known image decoders from the standard library.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	// Register all known image decoders from the extension library.
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
)

// LoadFile loads and decodes the named file, returning an image and the
// name of its format.  If an error occurs, the returned image is nil and
// the value of format is undefined.
func LoadFile(filename string) (image.Image, string, error) {
	rdr, err := os.Open(filename)
	if err != nil {
		return nil, "", err
	}
	defer rdr.Close()
	return image.Decode(rdr)
}

// Decode reads from an io.Reader, returning an image and the name of its
// format.  If an error occurs, the returned image is nil and the value of
// format is undefined.
func Decode(r io.Reader) (image.Image, string, error) {
	return image.Decode(r)
}
