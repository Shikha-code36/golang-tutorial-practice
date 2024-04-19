package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

/*
type gasEngine struct {
	mpg     uint8
	gallons uint8
ownerInfo owner
}

type owner struct{
	name string
}

and we can access in main function like this
var myEngine gasEngine = gasEngine{25, 50, owner{"Shikha"}}

can be accessed like 
myEngine.ownerInfo.name

but we are accessing field and having subfield and accessing the we can put direclty in struct

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner // later added
	int //later added 
}

*/

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

//methods of structs

func (e gasEngine) milesLeft() uint8{
	return e.gallons*e.mpg
}

func (e electricEngine) milesLeft() uint8{
	return e.kwh*e.mpkwh
}


/*
func canMakeIt (e gasEngine, miles uint8){
	if miles<=e.milesLeft(){
		fmt.Println("You can make it there!!!!!")
	}else{
		fmt.Println("Need to fuel up bro!!!!")
	}
}

currently canMakeIt function takes only for gasEngine what if we allow for all engine liek electricalEngine
and so we need it generic and then interfaces will help us
*/

// Let's define Interfaces

type engine interface{
	milesLeft() uint8
}

func canMakeIt (e engine, miles uint8){
	if miles<=e.milesLeft(){
		fmt.Println("You can make it there!!!!!")
	}else{
		fmt.Println("Need to fuel up bro!!!!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 50}

	//ways of defining fields 
	// var myEngine gasEngine = gasEngine{mpg:25, gallons:50}
	// or directly like:
	// myEngine.mpg = 25
	// myEngine.gallons = 50

	fmt.Printf("Original struct myengine mpg %v and gallons %v\n",myEngine.mpg, myEngine.gallons)

	// anonymous structs defined similiar to original gasEngine struct but cannot be reusable. That means we need to define similiar struct again to use.
	var myEngine2 = struct{
		mpg uint8
		gallons uint8
	}{12,11} // not reusable struct need to define similiar ex- myEngine3 and stuffs
	
	fmt.Printf("Anonymous struct myengine2 mpg %v and gallons %v\n",myEngine2.mpg, myEngine2.gallons)

	fmt.Printf("Total miles left in tank for OG: %v\n", myEngine.milesLeft())

	var myEngineElec electricEngine = electricEngine{25, 50}

	canMakeIt(myEngineElec, 35)	
}


/*
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> go run tutorial6/main.go

O/P :::::::::::::::::

Original struct myengine mpg 25 and gallons 50
Anonymous struct myengine2 mpg 12 and gallons 11
Total miles left in tank for OG: 226
You can make it there!!!!!

*/