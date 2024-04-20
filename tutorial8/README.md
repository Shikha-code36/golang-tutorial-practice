## Concurrency(Goroutines) in GO

Concurrency is the execution of multiple instruction sequences at the same time. In Go, concurrency is achieved using goroutines. A goroutine is a lightweight thread of execution managed by the Go runtime. You can launch a function as a goroutine by prefixing the function call with the `go` keyword.

```go
go dbCall(i)
```

However, if we run the code like this, it will feel like nothing happened. Our program spawns these tasks in the background, doesn't wait for them to finish, and then exits the program before they are complete. This is where wait groups come in.

## WaitGroup

A WaitGroup is used to wait for a collection of goroutines to finish executing. The main goroutine calls `Add` to set the number of goroutines to wait for. Then each of the goroutines runs and calls `Done` when finished. At the same time, `Wait` can be used to block until all goroutines have finished.

```go
var wg = sync.WaitGroup{}

func main(){
    for i:=0; i<len(dbData); i++{
        wg.Add(1)
        go dbCall(i)
    }
    wg.Wait()
}
```

## Mutex

A Mutex, or a mutual exclusion lock, is a synchronization primitive that can be used to protect shared data from being concurrently accessed by multiple goroutines. In Go, `sync.Mutex` is used for this purpose.

```go
var m = sync.Mutex{}

func save(result string){
    m.Lock()
    results = append(results, result)
    m.Unlock()
}
```

However, one drawback of a mutex is that it completely locks out other goroutines from accessing a result slice. This is where `sync.RWMutex` comes in.

## RWMutex

`sync.RWMutex` is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer. It has `RLock` and `RUnlock` methods for reading, and `Lock` and `Unlock` methods for writing.

```go
var m = sync.RWMutex{}

func save(result string){
    m.Lock()
    results = append(results, result)
    m.Unlock()
}

func log(){
    m.RLock()
    fmt.Printf("\nThe current results are: %v", results)
    m.RUnlock()
}
```

### Checkout the code

 - [main.go](main.go)

