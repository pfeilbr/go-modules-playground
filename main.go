package main

import (
	"fmt"

	"github.com/tutorialedge/go-modules-test/lib01"
	"github.com/tutorialedge/go-modules-test/lib02"
)

// EchoQuote echo and quote string
func EchoQuote(s string) string {
	return lib01.Echo(lib02.Quote("hello"))
}

func main() {
	fmt.Printf("lib01.Echo(%v): %v\n", "hello", lib01.Echo("hello"))
	fmt.Printf("lib01.Quote(%v): %v\n", "hello", lib02.Quote("hello"))
}
