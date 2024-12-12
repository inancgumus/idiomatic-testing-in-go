package main

import "fmt"

func main() {
	fileServer := &server{url: "file"}
	fileServer.check()
	fmt.Printf("is slow? %t\n", fileServer.slow())
}
