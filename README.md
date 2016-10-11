# ish

A collection of algorithms for comparing the similarity of images.

## Usage

This [example program](example/dhash.go) will compute the difference hash of a
list of images from the command-line arguments:

```go
package main

import (
	"encoding/hex"
	"fmt"
	"os"

	"github.com/brett-lempereur/ish"
)

func main() {
	hasher := ish.NewDifferenceHash(8, 8)
	for _, filename := range os.Args[1:] {
		img, ft, err := ish.LoadFile(filename)
		if err != nil {
			fmt.Printf("%s: %s\n", filename, err.Error())
			continue
		}
		dh, err := hasher.Hash(img)
		if err != nil {
			fmt.Printf("%s <%s>: %s", filename, ft, err)
		} else {
			dhs := hex.EncodeToString(dh)
			fmt.Printf("%s <%s>: %s\n", filename, ft, dhs)
		}
	}
}
```

Running this program against the test data in this repository gives:

```
[brett@saito ish]$ go run example/dhash.go testdata/*
testdata/house-32x32-Gray.png <png>: df8681cbc31986d9
testdata/house-4MP-CMYK.png <png>: df8681cbcb19861b
testdata/house-4MP-Gray.png <png>: df8681cbc3198699
testdata/house-4MP-RGBA.jpeg <jpeg>: df8681cbcb19865b
testdata/house-4MP-RGBA.png <png>: df8681cbcb19861b
testdata/house-8MP-CMYK.png <png>: df8681cbcb19861b
testdata/house-8MP-Gray.png <png>: df8681cbc3198699
testdata/house-8MP-RGBA.jpeg <jpeg>: df8681cbcb19865b
testdata/house-8MP-RGBA.png <png>: df8681cbcb19861b
```

## Installation

Use `go get` to pull the package and dependencies into your `GOPATH`:

    go get -u github.com/brett-lempereur/ish

## Thanks

Thanks to Dr. Neal Krawetz for documenting some efficient image similarity
hashing algorithms over at the [HackerFactor][hf].

## License

Released under the [BSD license](LICENSE.md).

[hf]: http://www.hackerfactor.com/blog/?/archives/529-Kind-of-Like-That.html
