package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		Error()
	}

	pieces := ReadPieces()
	BuildSolution(pieces)
}

func Error() {
	fmt.Println("ERROR")
	os.Exit(0)
}
