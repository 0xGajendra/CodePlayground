//Range and close
// a sender can close a channl to indicate that no more values will be sent. Receivers can test whther a channel has been closed by assigning a second paramter to the receive expression: after

//v, ok := <- ch

// ok is false if there are no more values to receive and the channel is closed

// The loop for i:= range c receives values from the channel repeatedly until it is closed

// Node: onlt the sender should close a channel, never the receiver. Sending a closed channl will cause a panic

//Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop

package main

import (
	"fmt"
)

func fibo(n int, c chan int){
	x,y := 0,1
	for i:=0; i<n; i++ {
		c<-x
		x,y = y, x+y
	}
	close(c)
}

func main(){
	c:= make(chan int, 10)
	go fibo(cap(c), c)
	for i:= range c {
		fmt.Println(i)
	}
}