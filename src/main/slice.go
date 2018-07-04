package main

import "fmt"

func modify2(array []int){
	array[0] = 10
	fmt.Println("In modify(), array values:", array)
}
