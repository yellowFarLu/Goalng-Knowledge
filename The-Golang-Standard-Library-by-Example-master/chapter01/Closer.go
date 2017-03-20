package main

import (
	"fmt"
	"os"
)

func main() {
	file, error := os.Open("dadadasdad.txt")
	if error != nil {
		fmt.Println(error)
	}

	defer file.Close()

}
