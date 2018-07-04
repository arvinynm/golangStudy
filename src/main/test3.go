package main

import (
	"fmt"
	"log"
)

type Rect struct {
	x, y float64
	width, height float64


}
func(r *Rect) Area() float64{
	return r.width * r.height
}

func main(){
	//r := Rect{
	//	x:1,
	//	y:2,
	//	width:4,
	//	height:5,
	//}
	r := new(Rect)
	fmt.Println(r.Area())
}

type foo struct {
	Rect
}
type Job struct {
	Command string
	log *log.Logger
}
func (job *Job)Start(){
	job.log.Flags()
}
