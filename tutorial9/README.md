# Go Channels 

This Go program demonstrates the use of channels for communication between goroutines.

## Key Concepts

### Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.

```go
var c = make(chan int) //this channel can only hold a single int value
c <- 1                 // adding value to this channel
var i = <-c            // retrieve the value from the channel now c will be empty and i = 1
fmt.Println(i)
```

If we run the above code directly, it will give a deadlock error. This is because the main goroutine is trying to receive a value from the channel, but there is no other goroutine available to send a value.

### Buffered Channels

Buffered channels can hold more than one value before they block. The capacity is passed as the second argument to `make`.

```go
var c = make(chan int, 5) // this channel can hold up to 5 int values
```

### Goroutines and Channels

Goroutines can send values into a channel, and other goroutines can receive those values.

```go
func main() {
    var c = make(chan int, 5)
    go process2(c)
    fmt.Println(<-c)
    for i:= range c{
        fmt.Println(i)
        time.Sleep(time.Second+1)
    }
}
```

In the `main` function, we create a buffered channel `c` and spawn a goroutine `process2(c)`. We then receive a value from the channel and print it. We also range over the channel to receive all remaining values.

### Closing Channels

Senders can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression.

```go
func process2(c chan int){
    defer close(c)
    for i:=0; i<5; i++{
        c <- i
    }
    fmt.Println("Exiting Process")
}
```

In the `process2` function, we defer the call to `close(c)` to ensure that the channel will be closed once all values have been sent. We then send five values into the channel and print a message indicating that the process is exiting.


### Checkout the code

- [main.go](main.go)