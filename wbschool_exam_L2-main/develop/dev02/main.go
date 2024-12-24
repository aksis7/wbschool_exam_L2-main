package main

import (
	"dev02/unpack"
	"fmt"
)

func main() {
	examples := []string{
		"a4bc2d5e",
		"qwe\\4\\5", // qwe45
		"qwe\\45",   // qwe44444
		"qwe\\\\5",  // qwe\\\\\
	}

	for _, example := range examples {
		result, err := unpack.Unpack(example)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("Input: %s => Output: %s\n", example, result)
	}
}
