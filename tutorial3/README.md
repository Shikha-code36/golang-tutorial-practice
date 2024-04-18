# Functions and Control Structures in Go

This will guide you through some basic functions and control structures in Go. The code we'll be discussing is a simple program that demonstrates the use of functions, if-else statements, and switch cases.

## Functions

### printMe Function

The `printMe` function is a simple function that prints a string to the console. Here's the function definition:

```go
func printMe(){
	fmt.Println("Hi! This is the function that we defined")
}
```

And here's how it's called in the `main` function:

```go
printMe()
```

When you run the program, it will print: "Hi! This is the function that we defined".

### printWithParam Function

The `printWithParam` function is similar to `printMe`, but it takes one parameter: a string that it will print to the console. Here's the function definition:

```go
func printWithParam(printValue string){
	fmt.Println(printValue)
}
```

And here's how it's called in the `main` function:

```go
var printValue string = "Parameter function is called"
printWithParam(printValue)
```

When you run the program, it will print: "Parameter function is called".

### intDivision Function

The `intDivision` function takes two integers as parameters and returns the result and remainder of integer division, along with an error if the denominator is zero. Here's the function definition:

```go
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
```

## Control Structures

### If-Else Statement

In Go, you can use if-else statements to execute different code blocks based on certain conditions. In our code, we use an if-else statement to handle the error returned by the `intDivision` function:

```go
var result, remainder, err  = intDivision(numerator,denominator)
if err!=nil{
	fmt.Print(err.Error())
}else if remainder == 0{
	fmt.Printf("The result of the integer division is %v", result)	
}else{
	fmt.Printf("The result of the integer division is %v with remainder %v",result, remainder)
}
```

### Switch Case

Switch cases in Go are a more readable alternative to if-else statements when you have multiple conditions to check. In our code, we use a switch case to print different messages based on the remainder of the division:

In our code, we use a switch case to print different messages based on the remainder of the division:

```go
switch remainder{
case 0:
	fmt.Println("The divison was exact")
case 1,2:
	fmt.Println("The division was close")
default:
	fmt.Println("The division was not close")
}
```

Here, `switch remainder` means we're switching on the value of `remainder`. 

- If `remainder` is `0`, we print "The divison was exact".
- If `remainder` is `1` or `2`, we print "The division was close".
- If `remainder` is any other value, we print "The division was not close". This is the `default` case, which is executed when none of the other cases match.

## Error Handling

In Go, errors are values that can be returned from functions. The `error` type is a built-in interface, and a `nil` error denotes success. Non-nil error denotes failure.

In our `intDivision` function, we return an error when the denominator is zero:

```go
if denominator == 0{
	err = errors.New("cannot divide by zero")
	return 0, 0 , err
}
```

Here, `errors.New("cannot divide by zero")` creates a new error with the message "cannot divide by zero". This error is then returned from the function.

In the `main` function, we check if `err` is not `nil`, which means an error occurred:

```go
if err!=nil{
	fmt.Print(err.Error())
}
```

If an error occurred, we print the error message using `err.Error()`. This is a method provided by the `error` interface that returns the error message.




### Checkout the code

 - [main.go](main.go)