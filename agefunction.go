package main

import (
	"fmt"
	"strconv"
)

func main() {
	age, err := getAge("25") // Pass an example age string here

	if err != nil {
		fmt.Println("Conversion Error:", err)
		return
	}

	fmt.Println("Age is:", age)
}

func getAge(ageStr string) (int, error) {
	ageAsInt, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, err
	}

	return ageAsInt, nil
}
