package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type contactInfo struct{
	Name string	
	Email string	
}
	
type purchaseInfo struct{	
	Name string
	Price float32
	Amount int
}

func main(){
	var intSlice = []int{1,2,3}
	fmt.Println(sumSlice[int](intSlice))

	var float32Slice = []float32{1,2,3}
	fmt.Println(sumSlice[float32](float32Slice))

	var sliceint = []int{}
	fmt.Println(isEmpty[int](sliceint))
	fmt.Println(isEmpty[int](intSlice))

	var contacts []contactInfo = LoadJSON[contactInfo]("./files/contactInfo.json")
	fmt.Printf("\n%+v", contacts)

	var purchases []purchaseInfo = LoadJSON [purchaseInfo]("./files/PurchaseInfo.json")
	fmt.Printf("\n%+v", purchases)
	}

func sumSlice[T int | float32 | float64](slice []T) T{
	var sum T
	for _, v := range slice{
		sum += v
	}
	return sum
}

// we can do type any because not all types are compaitible with + operators

func isEmpty[T any](slice []T) bool{
	return len(slice) == 0
}

func LoadJSON[T contactInfo | purchaseInfo] (filePath string) []T{
	data, _ := ioutil.ReadFile(filePath)	
	var loaded = []T{}	
	json.Unmarshal(data, &loaded)

	return  loaded
}

/*
>>>>>>>>>>>>>>>>>>>>>>> go run tutorial10/main.go

O/P :::::::::::::::::::

6
6
true
false

[]
[]
PS 
*/