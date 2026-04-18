package main

import (
	"fmt"
)
//a closure is a function that references variable from outside it's body. the function may access and assign to referenced variables
func adder() func (int) int {
	sum:= 0
	return func(x int) int {
		sum+=x
		return sum
	}
}

func main(){
	pos, neg := adder(), adder()
	for i:=0; i<10; i++{
		fmt.Println(
			pos(i),
			neg(-2*i))
	}
}