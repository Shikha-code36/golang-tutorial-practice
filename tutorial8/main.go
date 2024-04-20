package main

import (
	"fmt"
	"time"
	"sync"
)

//var m = sync.Mutex{} // we can use the mutex to control the writing to or slice in a way that makes it safe in a concurrent program
var m = sync.RWMutex{}
var wg = sync.WaitGroup{}

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		wg.Add(1) //whenever we spawn a new go routine need to make sure to add 1 to the counter
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v",results)
}

func dbCall(i int){
	// Simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("\nThe result from the database is:", dbData[i])
	// initial code during mutex
	// m.Lock()
	// results = append(results, dbData[i])
	// m.Unlock()
	save(dbData[i])
	log()
	wg.Done()
}

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
// one drawback of mutex is it completely locks out othe goroutines to accessing a result slice

/*
func main(){
	t0 := time.Now()
	for i:=0; i<len(dbData); i++{
		dbCall(i)
	}
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int){
	// Simulate DB call delay
	var delay float32 = rand.Float32()*2000
	time.Sleep(time.Duration(delay)*time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
}

this gave output 
The result from the database is: id1
The result from the database is: id2
The result from the database is: id3
The result from the database is: id4
The result from the database is: id5

Total execution time: 4.0624966s

Another way to do this good way to let these db calls run concurrently to do this we use 
'go' keyword infront of the function you want to run concurrently

go dbCall(i)

if we run code like this it will feel like that nothing happened
so our program spawn these tasks in the background, didn't wait for them to finish
and then exited the program before they were complete

we need a way for our program to wait until all tasks completed and continue on to the rest of the code
this where wait groups comes in, which can be imported through sync package
*/

/*
>>>>>>>>>>>>>>>>>>>>>>> go run tutorial8/main.go

O/P :::::::::::::::

The result from the database is: id1

The current results are: [id1]
The result from the database is: id5

The current results are: [id1 id5]
The result from the database is: id2

The current results are: [id1 id5 id2]
The result from the database is: id3

The current results are: [id1 id5 id2 id3]
The result from the database is: id4

The current results are: [id1 id5 id2 id3 id4]
Total execution time: 2.0148696s
The results are [id1 id5 id2 id3 id4]
*/