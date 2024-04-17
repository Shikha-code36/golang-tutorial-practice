package main

// Every program must start with the package declaration. In Go language, packages are used to organize and reuse the code. 
// In Go, there are two types of program available one is executable and another one is the library. 
// The executable programs are those programs that we can run directly from the terminal and Libraries are the package of programs that we can reuse them in our program. 
// Here, the package main tells the compiler that the package must compile in the executable program rather than a shared library. 
// It is the starting point of the program and also contains the main function in it.

import(
	"fmt"
	)

// Here, import keyword is used to import packages in your program and fmt package is used to implement formatted Input/Output with functions.

func main() {
    fmt.Println("!... Hello World ...!")
}

// In the above lines of code we have:
// func: It is used to create a function in Go language.
// main: It is the main function in Go language, which doesn’t contain the parameter, doesn’t return anything, and call when you execute your program.
// Println(): This method is present in fmt package and it is used to display “!… Hello World …!” string. Here, Println means Print line.