package main

import (
	"fmt"
	"math"
)

// I is an interface with one method: anything that has an M() method
// automatically satisfies this interface.
type I interface{
	M()
}

// T is a simple struct type that stores a string.
type T struct{
	S string
}


// M is the method that makes *T satisfy the I interface.
// The receiver is a pointer, so a value of type *T can call this method.
func (t *T) M(){
	fmt.Println(t.S)
}

// F is a custom type based on float64.
type F float64

// M is also defined for F, so F satisfies the I interface too.
func (f F) M() {
	fmt.Println(f)
}


// describe prints two things:
// 1. the value stored inside the interface
// 2. the concrete type of that value
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

// main demonstrates interface values in Go.
func main(){
	var i I

	// Store a *T value inside the interface.
	// Even though the variable type is I, the actual value is *T.
	i=&T{"Hello"}
	describe(i)
	i.M()

	// Now store an F value inside the same interface variable.
	// This works because F also has an M() method.
	i=F(math.Pi)

	describe(i)
	i.M()
}
