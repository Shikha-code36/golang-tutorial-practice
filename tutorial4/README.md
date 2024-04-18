# Arrays, Slices, and Control Structures in Go

This will guide you through some basic concepts in Go, including arrays, slices, and control structures. The code we'll be discussing demonstrates these concepts in a simple and understandable way.

## Arrays

In Go, an array is a numbered sequence of elements of a specific length. Here are some ways you can declare an array:

1. Declare an array without initializing elements (all elements get the zero value of the element type):

```go
var intArr [3]int32
fmt.Println(intArr[0]) // prints: 0
```

2. Declare an array and initialize some or all elements:

```go
var intArr [3]int32
intArr[1] = 123
fmt.Println(intArr[0]) // prints: 0
fmt.Println(intArr[1]) // prints: 123
```

3. Declare and initialize an array using a literal:

```go
var intArr [3]int32 = [3]int32{1,2,3}
// or
intArr := [3]int32{1,2,3}
// or let Go figure out the length based on the number of elements
intArr := [...]int32{1,2,3}
```

4. Access a range of elements (a slice of the array):

```go
fmt.Println(intArr[1:3]) // prints: [2 3]
```

## Slices

A slice is a segment of an array. Slices are indexable and have a length. Unlike arrays, they can be resized. Here's how you can declare a slice:

```go
var intSlice []int32 = []int32{4,5,6}
```

You can append elements to a slice using the `append` function:

```go
intSlice = append(intSlice, 7)
```

You can also append another slice to a slice:

```go
var intSlice2 []int32 = []int32{8,9}
intSlice = append(intSlice, intSlice2...)
```

The `...` is used to pass the elements of `intSlice2` as separate arguments to the `append` function.

## Maps

A map is an unordered collection of key-value pairs. Here's how you can declare a map:

```go
var myMap map[string]uint8 = make(map[string]uint8)
```

You can add elements to the map like this:

```go
var myMap2 =  map[string]uint8{"Shikha":24, "Sarah":23, "Kanha":1}
```

## Control Structures

### For Loop

Go has only one looping construct: the `for` loop. The basic `for` loop has three components separated by semicolons: the init statement, the condition expression, and the post statement.

Here's a `for` loop that prints the numbers from 0 to 9:

```go
for i:=0; i<10; i++ {
	fmt.Printf("printing number %v \n",i)
}
```

### While Loop

Go doesn't have a `while` loop, but you can create one using the `for` loop:

```go
var i int = 0
for i<10{
	fmt.Println(i)
	i=i+1
}
```

This will print the numbers from 0 to 9, just like the previous `for` loop.


### Checkout the code

 - [main.go](main.go)