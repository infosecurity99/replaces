package main

import "fmt"

func main() {
	name := getNamePointer("jhon doe")
	fmt.Println("Name as a pointer", name)
	fmt.Println("Name is a text", *name)
}

func getNamePointer(name string) *string {
	return &name
}
