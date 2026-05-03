package main

//a go routine is a lightwight thread managed by the go runtime

// go f(x,y,z)

// starts with a new goroutine running
// f(x,y,z)

//the evaluation of f,x,y and z happens in the current goroutine and the execution of f happens in the new goroutine.
//go routines run in the same address space, so access to shared memory must be synchronied. The sync package provides useful primitives, although you won't need them in GO as there are other primitives.

import (
	"fmt"
	"time"
)
func say(s string){
	for i:= 0; i<5; i++ {
		time.Sleep(100*time.Millisecond)
		fmt.Println(s)
	}
}

func main(){
	go say("world")
	say("hello")
}