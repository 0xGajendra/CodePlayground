//Channels are a typed conduit through which you can send and receive value with the channel operator, <-.

// ch <- v //save v to channel ch.
// v := c<- ch //recieve from ch, and assign value to v

// the data flows in the direction of the arrow

// like the maps and slices, channels must be created before use

// ch:=make(chan int)

// by default, sned and receives block until the other side is ready.
// this allows goroutines to synchronize without explicit locks or condition variables

// the example code sums the numbers in a slice, disturbing the work between two goroutines. Once both goroutines have completed their computation their computation, it calculates the final result.

package main

import "fmt"

func sum(s []int, c chan int){
	sum:=0
	for _, v := range s {
		sum+=v
	}
	c<-sum //send sum to c
}

func main(){
	s:= []int {7,2,8,-9,4,0}

	c:= make(chan int)
	go sum(s[:len(s)/2],c)
	go sum(s[len(s)/2:], c)

	x,y := <- c, <-c //receive from c

	fmt.Println(x,y,x+y)
}