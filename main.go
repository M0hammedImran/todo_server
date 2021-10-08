package main

import (
	"fmt"
	"os"
)

func main() {
	os := os.Getenv("USERNAME")
	
	fmt.Printf("Hello world %s \n", os)
}