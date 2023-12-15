package main

import (
	"fmt"
	"slices"
)

func main() {
	buffer := []byte{
		137, 80, 78, 71, 13, 10, 26, 10, 65,
		/* ..more.. */
	}
	fmt.Println(isPNG(buffer)) // prints true
}

var pngHeader = [8]byte{ //nolint:gochecknoglobals
	137, 80, 78, 71, 13, 10, 26, 10,
}

func isPNG(src []byte) bool {
	if len(src) < len(pngHeader) {
		return false
	}
	return slices.Equal(src[:len(pngHeader)], pngHeader[:])
}
