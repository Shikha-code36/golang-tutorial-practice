package main

import (
	"errors"
	"fmt"
)

func main(){
	printMe()
	var printValue string = "Parameter function is called"
	printWithParam(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err  = intDivision(numerator,denominator)
	// if err!=nil{
	// 	fmt.Print(err.Error())
	// }else if remainder == 0{
	// 	fmt.Printf("The result of the integer division is %v", result)	
	// }else{
	// 	fmt.Printf("The result of the integer division is %v with remainder %v",result, remainder)
	// }
	switch{
	case err!=nil:
		fmt.Print(err.Error())
	case remainder==0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v",result, remainder)
	}

	switch remainder{
	case 0:
		fmt.Println("The divison was exact")
	case 1,2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}
}
	
func printMe(){
	fmt.Println("Hi! This is the function that we defined")
}

func printWithParam(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0{
		err = errors.New("cannot divide by zero")
		return 0, 0 , err
	}
	var result int = numerator/denominator
	var remainder int = numerator%denominator
	return result, remainder, err
}

/*
>>>>>>>>>>> go run tutorial3/main.go

O/p::::

Hi! This is the function that we defined
Parameter function is called
The result of the integer division is 5 with remainder 1The division was close

*/