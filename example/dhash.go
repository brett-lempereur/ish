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
