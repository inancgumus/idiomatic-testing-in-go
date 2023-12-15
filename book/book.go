// Package book offers information about my book.
package book

// Title returns the title of this book.
func Title() string {
	return "Go by Example: " + subtitle()
}

func subtitle() string {
	return "Programmer's Guide to Idiomatic and Testable Code"
}
