
// Go does not have classes, but you can define methods on types.
// A method is a function with a special "receiver" argument.

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// Abs returns the Euclidean distance from the origin to the vertex.
func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// func Abs(v Vertex) float64 {
// 	return math.Sqrt(v.X*v.X + v.Y*v.Y)
// }

//the above is the same 

func main() {
    v := Vertex{3, 4} // Create a Vertex with X=3, Y=4
    fmt.Println(v.Abs()) // Print the distance from origin
}