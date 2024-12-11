package main

import "fmt"

func main() {
	fileServer := server{url: "file"}
	fmt.Printf(
		"url: %s response time: %d\n",
		fileServer.url, fileServer.responseTime,
	)
}
