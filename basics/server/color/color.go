package main

import "fmt"

// Additional example for types with value receivers
// This one uses bitwise operators to invert the color.

type color int

func (c color) r() int { return int((c >> 16) & 0xFF) }
func (c color) g() int { return int((c >> 8) & 0xFF) }
func (c color) b() int { return int(c & 0xFF) }

// 1.
// func (c color) invert() { c = ^c & 0xFFFFFF }
//
// 2.
func (c color) invert() color { return ^c & 0xFFFFFF }

func main() {
	c := color(0x29BEB0)
	fmt.Println(c.r(), c.g(), c.b()) // prints 41 190 176

	// 1.
	//
	// can't mutate c because it's a value receiver
	// c.invert()
	// fmt.Println(c.r(), c.g(), c.b()) // still prints 41 190 176

	// 2.
	c = c.invert()
	fmt.Println(c.r(), c.g(), c.b()) // prints 214 65 79
}
