package main

import "fmt"
import "unicode/utf8"

func  main() {
	var intNum int = 35
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var intNum1 int = 6
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	fmt.Println(intNum1%intNum2)

	var myString string = "RadhaRani"
	fmt.Println(myString)

	// we can use back quotes for new line of string or 
	// we can use /n Radha /nRani 

	var mys1 string = `Radha
	Rani`
	fmt.Println(mys1)

	//we can also cancatenate string by adding them

	var myStr string = "Hello" + " " + "World"
	fmt.Println(myStr)

	/* if we use len to check length of string it will give no of bytes not no of characters
	since go uses UTF-8 encoding. Characters outside the ASCII characters set
	are stored with more than a single byte.

	getting length of a character se "S" it will give 1 but suppose we have special character like "γ"(gamma)
	we will use built in package unicode/utf-8
	*/

	fmt.Println(utf8.RuneCountInString("γ"))
	// O/P - 1

	// rune is also a data type in GO

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = false
	fmt.Println(myBoolean)  

	// var myVar = 'text'
	myVar := "text"
	fmt.Println(myVar)

	// var var1, var2 int = 1, 2
	var1, var2 := 1,2
	fmt.Println(var1,var2)

	const myConst string = "const value"
	fmt.Println(myConst)

	const pi float32 = 3.1415
	fmt.Println(pi)
}

/*

------ go run tutorial2/main.go

O/P ---

35
1.23456789e+07
3
0
RadhaRani
Radha
        Rani
Hello World
1
97
false
text
1 2
const value
3.1415
*/