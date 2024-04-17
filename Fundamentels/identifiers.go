package main

import "fmt"

func main(){
	var name = "SiyaRaam"
	fmt.Println("Jai", name)
}

/*
There is a total of three identifiers available in the above example:
	- main: Name of the package
	- main: Name of the function
	- name: Name of the variable
*/

/*
Rules for Defining Identifiers:
	1. The name of the identifier must begin with a letter or an underscore(_). And the names may contain the letters ‘a-z’ or ’A-Z’ or digits 0-9 as well as the character ‘_’.
	2. The name of the identifier should not start with a digit.
	3. The name of the identifier is case-sensitive.
	4. Keywords are not allowed to be used as an identifier name.
	5. There is no limit on the length of the name of the identifier, but it is advisable to use an optimum length of 4 – 15 letters only.
*/

/*
Example::::

// Valid identifiers:
_geeks23
geeks
gek23sd
Geeks
geeKs
geeks_geeks

// Invalid identifiers:
212geeks
if
default

*/

/*
In the Go language, there are some predeclared identifiers available for constants, types, and functions. 
These names are not reserved, you are allowed to use them in the declaration. 
Following is the list of predeclared identifiers:

For Constants:
true, false, iota, nil

For Types:
int, int8, int16, int32, int64, uint,
uint8, uint16, uint32, uint64, uintptr,
float32, float64, complex128, complex64,
bool, byte, rune, string, error

For Functions:
make, len, cap, new, append, copy, close, 
delete, complex, real, imag, panic, recover
*/

/*
-- The identifier represented by the underscore character(_) is known as a blank identifier. It is used as an anonymous placeholder instead of a regular identifier, and it has a special meaning in declarations, as an operand, and in assignments.
-- The identifier which is allowed to access it from another package is known as the exported identifier. The exported identifiers are those identifiers which obey the following conditions:
	--- The first character of the exported identifier’s name should be in the Unicode upper case letter.
	--- The identifier should be declared in the package block or be a variable, function, type, or method name within that package.
*/

/*

In the below example, file1.go contains an exported variable called ExportedVariable, which is accessible within the same file. 
It also imports the file2 package and accesses the exported variable AnotherExportedVariable from file2.go. By running go run file1.go, it will print the value of ExportedVariable (“Hello, World!”) from file1.go and the value of AnotherExportedVariable (“Greetings from file2!”) from file2.go. 
This demonstrates the concept of exported identifiers being accessible from another package in Go.
*/

/*
//file2.go

package file2

// Exported variable
var AnotherExportedVariable = "Greetings from file2!"
*/

/*
// file1.go

package main

import (
	"fmt"
	"github.com/yourusername/project/file2"
)

// Exported variable
var ExportedVariable = "Hello, World!"

func main() {
	// Accessing exported identifier in the same file
	fmt.Println(ExportedVariable)

	// Accessing exported identifier from another package
	fmt.Println(file2.AnotherExportedVariable)
}
*/

/*
Output: (after running go run file1.go)

Hello, World!
Greetings from file2!
*/
