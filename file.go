package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.log")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	buf := make([]byte, 1024)

	for {
		n, err := file.Read(buf)

		if err == io.EOF {
			fmt.Println("Done reading data.log file")
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if n > 0 {
			fmt.Println(string(buf[:n]))
		}
	}
}
