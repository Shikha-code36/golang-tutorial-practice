package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[1]
	fmt.Println(indexed) //if we run this we get a number 

	//let's print out the type of the index value
	fmt.Printf("%v %T \n", indexed, indexed) //even wierder we get unsigned int8

	for i,v := range myString{
		fmt.Println(i,v)
	} // we get index start with 0 and bunch of numbers with it's value which is bytes due to utf-8

	fmt.Printf("\n The length of 'myString' is %v", len(myString))

	// easier way to deal with the strings for indexing and iterating is to cast them in array of runes
	
	fmt.Println("\n----------After rune------------")
	var myString2 = []rune("résumé")
	var indexed2 = myString2[1]
	fmt.Println(indexed2) //if we run this we get a number 

	//let's print out the type of the index value
	fmt.Printf("%v %T \n", indexed2, indexed2) //now we get int32

	for i,v := range myString2{
		fmt.Println(i,v)
	} // we get index start with 0 and bunch of numbers with it's value which is bytes due to utf-8

	fmt.Printf("\n The length of 'myString2' is %v", len(myString2))

	// we can concantenate string using "+" symbol

	var strSlice = []string{"s", "h", "i", "k", "h", "a"}
	var catStr = ""
	for i := range strSlice{
		catStr += strSlice[i]
	}
	fmt.Printf("\nString building: %v",catStr)

	// strings are immutable in GO we are creating a new string catStr
	// we can use built in package for string building

	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	var catStr1 = strBuilder.String()
	fmt.Printf("\nString building through built in package: %v",catStr1)

}

/*
>>>>>>>>>>>>>>>>>>>>>>>>>>>> go run tutorial5/main.go

O/P :::::::::::::::::::::

195
195 uint8 
0 114
1 233
3 115
4 117
5 109
6 233

 The length of 'myString' is 8
----------After rune------------
233
233 int32 
0 114
1 233
2 115
3 117
4 109
5 233

 The length of 'myString2' is 6
String building: shikha
String building through built in package: shikha
*/