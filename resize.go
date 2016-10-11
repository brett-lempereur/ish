package ish

import (
	"image"

	"github.com/bamiaux/rez"
)

// Resize resizes a grayscale image using an efficient bicubic filter on a
// single thread.
func Resize(img *image.Gray, width, height int) (*image.Gray, error) {

	var err error

	// The resizing algorithm performs extremely poorly for target dimensions
	// less than 16x16.  Computing the resize in two-stages, first to 16x16 and
	// then to the target dimensions works around this.
	if width < 16 && height < 16 {
		img, err = Resize(img, 16, 16)
		if err != nil {
			return nil, err
		}
	}

	// Convert the input image to the target dimensions using a converter
	// configured to run on a single thread.
	output := image.NewGray(image.Rect(0, 0, width, height))
	cfg, err := rez.PrepareConversion(output, img)
	if err != nil {
		return nil, err
	}
	cfg.Threads = 1
	converter, err := rez.NewConverter(cfg, rez.NewBicubicFilter())
	if err != nil {
		return nil, err
	}
	err = converter.Convert(output, img)
	return output, err

}
