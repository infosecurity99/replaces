package main

import "fmt"

func main() {
	printTechStack("Go", "Gin", "gRPC", "Docker", "Kubernets", "AWS")
}

func printTechStack(names ...string) {
	for _, techName := range names {
		fmt.Printf("===>%s \n ", techName)
	}
}
