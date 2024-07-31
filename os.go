package main

import (
	"fmt"
	"os"
)

func FullName(first, last string) string {
	return fmt.Sprintf("%s %s", first, last)
}

func main() {
	first := os.Args[1]
	last := os.Args[2]

	fmt.Printf("FirstName is %s, lastName is %s. Your Fullname is %s", first, last, FullName(first, last))
}
