package main

func main() {
	selfie := []byte{ /* ~100 MB pixel data */ }
	invertColors(selfie)
}

func invertColors(colors []byte) {
	for i := range colors {
		colors[i] = 255 - colors[i]
	}
}
