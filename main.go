package main

import (
	"fmt"
	"go-reloaded/processor"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	result := processor.Process(string(data))

	err = os.WriteFile(outputFile, []byte(result+"\n"), 0644)
	if err != nil {
		fmt.Println("Error writing output file:", err)
	}
}
