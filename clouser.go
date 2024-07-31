package main

import (
	"fmt"
	"sort"
)

func main() {
	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Amit", "Kumar", 42},
		{"Ram", "Singh", 18},
		{"Tulip", "Sharma", 51},
	}
	fmt.Println(people)

	// sort by last name
	sort.Slice(people, func(i int, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println(people)
}
