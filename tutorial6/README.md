# Structs and Interfaces in Go

## Structs

A `struct` is a collection of fields. It's a way to group together related data into a single unit. 

### Defining a Struct

You can define a struct like this:

```go
type gasEngine struct {
	mpg     uint8
	gallons uint8
}
```

### Ways of Defining Fields

There are several ways to define fields in a struct:

1. Directly assign values in order:

```go
var myEngine gasEngine = gasEngine{25, 50}
```

2. Assign values by field name:

```go
var myEngine gasEngine = gasEngine{mpg:25, gallons:50}
```

3. Assign values one by one:

```go
myEngine.mpg = 25
myEngine.gallons = 50
```

### Nested Structs and Anonymous Fields

Structs can also contain other structs as fields. This is useful when you want to group related data together. For example:

```go
type owner struct{
	name string
}

type gasEngine struct {
	mpg     uint8
	gallons uint8
	ownerInfo owner
}
```

You can then create a `gasEngine` with an `owner` like this:

```go
var myEngine gasEngine = gasEngine{25, 50, owner{"Shikha"}}
```

And access the `name` of the `owner` like this:

```go
myEngine.ownerInfo.name
```

This is known as a nested struct because the `owner` struct is nested inside the `gasEngine` struct.

#### Anonymous Fields

Go also supports anonymous fields. If you have a field in your struct without a name, just the type, Go will use the type as the name of the field.

```go
type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner // anonymous field
	int // anonymous field
}
```

In this case, `owner` and `int` are anonymous fields. You can access them directly:

```go
myEngine.owner
myEngine.int
```

### Anonymous Structs

Anonymous structs are similar to regular structs but they are not reusable. You need to define a similar struct again to use it.

```go
var myEngine2 = struct{
	mpg uint8
	gallons uint8
}{12,11}
```

## Interfaces

Interfaces define a method set called its interface. Any other type is said to implement this interface if it provides definitions for all the methods in the interface.

### Defining an Interface

Here's how you can define an interface:

```go
type engine interface{
	milesLeft() uint8
}
```

This `engine` interface can be implemented by any struct that has a `milesLeft()` method returning a `uint8`.

### Using Interfaces

Interfaces allow us to write functions that can take different types of arguments. For example, the `canMakeIt` function can take any type that implements the `engine` interface:

```go
func canMakeIt (e engine, miles uint8){
	if miles<=e.milesLeft(){
		fmt.Println("You can make it there!!!!!")
	}else{
		fmt.Println("Need to fuel up bro!!!!")
	}
}
```

This function can now work with any type of engine, not just `gasEngine`.

## Conclusion

Structs and interfaces are powerful features in Go that allow you to group related data and define behavior for different types. They are fundamental to writing clear and maintainable code in Go.


### Checkout the code

 - [main.go](main.go)