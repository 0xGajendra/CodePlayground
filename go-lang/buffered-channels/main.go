//Channels can be buffered. Provide the buffer length as the second argument to ``make`` to initialize a buffered channel:

// syntax

// ch:=make(chan int, 100)

// sends to a buffered channel block only when only when the buffer is full.

// receives block when the buffer is empty

package main

import "fmt"

func main(){
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	ch <- 3
	ch<-4
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
