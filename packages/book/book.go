// Package book offers information about the Go by Example book.
package book

// Title returns the title of this book.
func Title() string {
	return "Go by Example: " + subtitle()
}

// subtitle returns the subtitle of this book.
func subtitle() string {
	return "Programmer's Guide to Idiomatic and Testable Code"
}
