package main

import "fmt"

func main() {
	//var p *int32 ---
	//variable p will hold memory address of an int32
	// but this pointer doesn't point to anything yet it's value is nil yet
	// in other words this pointer does not have an address assigned to it

	// to give pointer an address we will use builtin new function

	var p *int32 = new(int32) //now p stores a memory location
	var i int32

	fmt.Printf("The value p points to is: %v" , *p)
	fmt.Printf("\nThe value if i is: %v",i) // initially they both will be zero value no value assigned

	// to change the values stored at memory location
	// make sure your pointer isn't nil before assigning
	// *p = 10

	// we can also create a pointer from the address of another varialble using '&'
	// p = &i now they both reference the same int32 value in memory

	p = &i
	*p = 1 // here value of i will also changes
	fmt.Printf("\nThe value p points after to is: %v" , *p)
	fmt.Printf("\nThe value  after if i is: %v",i)

	/*
	If we use with regular variable
	var k int32 = 2
	i = k //the value of i also changes
	
	The main exception of this copy behaviour of non-pointer variables is when working with slices
	*/
	var slice = []int32{1,2,3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Printf("\nOG slice: %v", slice)
	fmt.Printf("\ncopy of slice: %v", sliceCopy)
	/*
	in terminal we will see value of original slice also changes
	because under the hood slices contains pointer to underlying array
	*/

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p",&thing1)
	//var result [5]float64 = square(thing1) this in case of copy
	var result [5]float64 = square(&thing1) 
	fmt.Printf("\nThe result is: %v",result)
	fmt.Printf("\nThe value of thing1 is: %v", thing1)
}	

func square(thing2 *[5]float64) [5]float64{
	fmt.Printf("\nThe memory location of the thing2 array is: %p",&thing2)
	for i := range thing2{
		thing2[i] =  thing2[i]*thing2[i]
	}
	return *thing2
}

/*
func square(thing2 [5]float64) [5]float64{
	fmt.Printf("\nThe memory location of the thing2 array is: %p",&thing2)
	for i := range thing2{
		thing2[i] =  thing2[i]*thing2[i]
	}
	return thing2
}
the memory addresses of thing1 and thing2 are different means these are 2 different arrays
that means we can modify the values of thing2 without effecting the thing1

currently in square we are doubling the memory usage, because we are creating a copy for use in function
we using way more memory then we need instead we will be using pointers
we can make our function taking a pointer to an array instead the memory location will be same
*/


/*
>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>> go run tutorial7/main.go

O/P ::::::::::::::::::::::::

The value p points to is: 0
The value if i is: 0
The value p points after to is: 1
The value  after if i is: 1
OG slice: [1 2 4]
copy of slice: [1 2 4]
The memory location of the thing1 array is: 0xc00000e450
The memory location of the thing2 array is: 0xc00005a030
The result is: [1 4 9 16 25]
The value of thing1 is: [1 4 9 16 25]
*/