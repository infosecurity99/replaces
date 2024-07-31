package main

import "fmt"

func main() {
	username, age := getUserDetails()

	fmt.Println("Username is :", username)
	fmt.Println("Age is:", age)
}

func getUserDetails() (string, int) {
	username := "JhonFoe"
	age := 34
	return username, age
}
