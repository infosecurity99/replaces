package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 4 {
			break
		}
		fmt.Println(i)
	}

	fmt.Println("Done")
}
