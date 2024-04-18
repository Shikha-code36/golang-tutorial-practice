package main

import "fmt"

func main() {
	// var intArr [3]int32
	// intArr[1] = 123
	// fmt.Println((intArr[0]))
	// fmt.Println(intArr[1:3])

	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	//var intArr [3]int32 = [3]int32{1,2,3}
	//intArr := [3]int32{1,2,3}
	intArr := [...]int32{1,2,3}
	fmt.Println(intArr)

	var intSlice []int32 = []int32{4,5,6}
	fmt.Printf("The length is %v with capacity %v and the value is %v", len(intSlice), cap(intSlice), intSlice)
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %v with capacity %v and the value is %v", len(intSlice), cap(intSlice), intSlice)

	var intSlice2 []int32 = []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Printf("\n for slice 2 %v", intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	intSlice2 = append(intSlice2, intSlice3...)
	fmt.Printf("\n for slice 3 %v \n", intSlice2)

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var myMap2 =  map[string]uint8{"Shikha":24, "Sarah":23, "Kanha":1}
	fmt.Printf("Age of Shikha %v \n",myMap2["Shikha"])

	var age, ok = myMap2["Adam"]
	if ok{
		fmt.Printf("The age is %v", age)
	}else{
		fmt.Println("Invalid Name")
	}

	delete(myMap2, "Sarah")

	for name, age:= range myMap2{
		fmt.Printf("Name: %v, Age:%v \n", name, age)
	}

	for i , v := range intArr{
		fmt.Printf("Index: %v, Value: %v\n", i, v)
	}

	//GO doesn't have while loop but it can be achieved by for loop itself

	var i int = 0
	for i<10{
		fmt.Println(i)
		i=i+1
	}

	for i:=0; i<10; i++ {
		fmt.Printf("printing number %v \n",i)
	}

}

/*

>>>>>>>>>>>>>>>>> go run tutorial4/main.go

O/P ::::::


[1 2 3]
The length is 3 with capacity 3 and the value is [4 5 6]
The length is 4 with capacity 6 and the value is [4 5 6 7]
 for slice 2 [4 5 6 7 8 9]
 for slice 3 [8 9 0 0 0] 
map[]
Age of Shikha 24 
Invalid Name
Name: Shikha, Age:24 
Name: Kanha, Age:1 
Index: 0, Value: 1
Index: 1, Value: 2
Index: 2, Value: 3
0
1
2
3
4
5
6
7
8
9
printing number 0
printing number 1
printing number 2
printing number 3
printing number 4
printing number 5
printing number 6
printing number 7
printing number 8
printing number 9
*/