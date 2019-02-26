package main

import (
	"fmt"
	"day2/leak"
)

func main() {

	fmt.Printf("***** Go routines leaks *****")
	fmt.Printf("The quickest response is: %s", leak.MirroredQuery())
}