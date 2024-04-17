package main

import (
	"fmt"
	"github.com/Shikha-code36/golang-tutorial-practice/Fundamentels/identifiers/file2/file2.go"
)

// Exported variable
var ExportedVariable = "Radhe Radhe!"

func main() {
	// Accessing exported identifier in the same file
	fmt.Println(ExportedVariable)

	// Accessing exported identifier from another package
	fmt.Println(file2.AnotherExportedVariable)
}