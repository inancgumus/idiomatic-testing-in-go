package main

import (
	"fmt"
	"time"
)

// --------------------------------------------------------------------------------
// Each main function contains the code for each section of the chapter.
//
// For example, think of mainForPointerReceivers() as the main function for the
// pointer receivers section, as described in the book.
// --------------------------------------------------------------------------------

func main() {
	// Comment out and uncomment the following function calls to see
	// the output of each section separately.

	mainForPointerReceivers()

	mainForValueReceivers()

	mainForInterfaces()
}

// mainForPointerReceivers is the main function for the
// pointer receivers section.
// this is called from main() in main.go.
func mainForPointerReceivers() { // func main() {
	// Creating a value of type server

	// fileServer := server{url: "file"}
	// fmt.Printf(
	// 	"url: %s response time: %d\n",
	// 	fileServer.url,
	// 	fileServer.responseTime,
	// )

	// Creating a pointer to a value of type server
	fileServer := &server{url: "file"}
	fileServer.check()
	fmt.Printf("is slow? %t\n", fileServer.slow())
}

// --------------------------------------------------------------------------------

// consider this func as the main function for the value receivers
// section of the book. this is called from main() in main.go.
func mainForValueReceivers() { // func main() {
	cpu := usage(0.99)
	fmt.Println("CPU usage:", cpu)             // prints 0.99
	fmt.Println("high CPU usage?", cpu.high()) // prints ..true

	cpu = cpu.set(0.7)
	fmt.Println("CPU usage:", cpu)         // prints 0.7
	fmt.Println("high usage?", cpu.high()) // prints ..false

	// incorrect:
	// cpu.set(0.5)                               // cpu is still 0.95 instead of 0.5
	// fmt.Println("high CPU usage?", cpu.high()) // still prints true
}

// --------------------------------------------------------------------------------

// mainForInterfaces is the main function for the interfaces section.
// this is called from main() in main.go.
func mainForInterfaces() { // func main() {
	// Comment out and uncomment the following function calls to see
	// the output of each section separately.
	mainForInterfacesIntro()
	mainForInterfacesMultiNotifier()
}

// --------------------------------------------------------------------------------

func mainForInterfacesIntro() { // func main() {
	authServer := &server{
		url:          "auth",
		responseTime: time.Minute,
	}

	slack := &slackNotifier{
		/* Slack specific configuration */
	}
	sms := &smsNotifier{
		/* SMS specific configuration   */
	}

	notify(authServer, slack) // prints: slack: auth server is slow: 1m0s
	notify(authServer, sms)   // prints: sms: auth server is slow: 1m0s
}

// --------------------------------------------------------------------------------

func mainForInterfacesMultiNotifier() { // func main() {
	authServer := &server{
		url:          "auth",
		responseTime: time.Minute,
	}

	slack := &slackNotifier{
		/* Slack specific configuration */
	}
	sms := &smsNotifier{
		/* SMS specific configuration   */
	}

	notify(authServer, multiNotifier{slack, sms})
}
