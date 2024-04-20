package main

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main(){
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"swiggy", "zomato", "eatclub"}
	for i:= range websites{
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
}

func checkChickenPrices(website string, chickenChannel chan string){
	for {
		time.Sleep(time.Second+1)
		var chickenPrice = rand.Float32()*20
		if chickenPrice<=MAX_CHICKEN_PRICE{
			chickenChannel <- website
			break
		}
	}
}

func checkTofuPrices(website string, tofuChannel chan string){
	for {
		time.Sleep(time.Second+1)
		var tofuPrice = rand.Float32()*20
		if tofuPrice<=MAX_TOFU_PRICE{
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string){
	select{
	case website := <-chickenChannel:
		fmt.Printf("\nText Sent: Found deal on chicken at %v.", website)
	case website := <-tofuChannel:
		fmt.Printf("Email Sent: Found deal on tofu at %v.", website)
	}
}



// func main() {
// 	//var c = make(chan int)
// 	var c = make(chan int, 5) //by adding this process function will run immediately but main function will take time
// 	//go process(c)
// 	go process2(c)
// 	fmt.Println(<-c)
// 	for i:= range c{
// 		fmt.Println(i)
// 		time.Sleep(time.Second+1)
// 	}
// }

// // func process(c chan int){
// // 	c <- 123
// // }

// func process2(c chan int){
// 	defer close(c)
// 	for i:=0; i<5; i++{
// 		c <- i
// 	}
// 	fmt.Println("Exiting Process")
// }

/*
var c = make(chan int) //this channel can only hold a single int value
c <- 1                 // adding value to this channel
var i = <-c            // retrieve the value from the channel now c will be empty and i = 1
fmt.Println(i)

if we run the above code direclty it will give deadlock error
*/
