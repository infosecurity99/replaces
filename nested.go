package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j > 0; j-- {

			if j == i {
				fmt.Println("Outer loop %d:Inner loop %d", i, i)
				break
			}
		}
	}
	fmt.Println("Done")
}
