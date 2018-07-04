package main

import "fmt"

var ch1 chan int
var ch2 chan<- float64
var ch3 <-chan int

func mai(){
	ch4 := make(chan int)
	ch5 := <-chan int(ch4)
	ch6 := chan<- int(ch4)
}

func Parse(ch <-chan int){
	for value := range ch{
		fmt.Println(value)
	}
	x, ok := <-ch
}

