# Pointers in Go

A pointer is a variable that stores the memory address of another variable. In Go, a pointer is represented using the `*` (asterisk) followed by the type of the stored value.

## Creating Pointers

You can create a pointer using the built-in `new` function:

```go
var p *int32 = new(int32) // p now stores a memory location
```

This creates a new `int32` value, sets it to zero (the zero value for integers), and returns a pointer to it.

## Using Pointers

You can change the value stored at the memory location a pointer points to:

```go
*p = 1 // changes the value at the memory location p points to
```

You can also create a pointer from the address of another variable using the `&` operator:

```go
p = &i // p and i now reference the same int32 value in memory
```

## Pointers and Arrays

When working with arrays, you can pass a pointer to the array to a function if you want the function to be able to modify the original array:

```go
func square(thing2 *[5]float64) [5]float64{
	for i := range thing2{
		thing2[i] =  thing2[i]*thing2[i]
	}
	return *thing2
}
```

In this function, `thing2` is a pointer to an array of `float64`. The function squares each element of the array. Because we're passing a pointer to the array, the original array is modified.

### Checkout the code

 - [main.go](main.go)