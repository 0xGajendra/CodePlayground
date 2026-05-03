package main

import "fmt"

// go functions are written to work on multiple types using type paramters. The type paramters of a function appear nbetween brackets, before the function's arguments

// syntax =  func Index[T comparable](s []T, x T) int

//Index returns the index of x in s, or -1 if not found

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constrainst, so we can use == here
		if v == x{
			return i
		}
	}
	return -1
}

//List represents a  single-linked list that holds
// values of any type

type List[T any] struct{
	next *List[T]
	val T
}


func main(){
	//Index works on a slice of ints
	si:=[]int {10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	//Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}