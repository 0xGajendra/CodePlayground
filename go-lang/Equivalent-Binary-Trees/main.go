/*
exercise: quivalent binary trees

1. Implement the walk function
2. Test the walk function

the function tree.New(k) constructs a randomly-structured(but always sorted) binary tree holding the values k, 2k, 3k,.. 10k.

create a new channel ch and kick off the walker
create a new channel ch and kich off the walker

go Walk(tree.New(1), ch)

then read and print 10 values from teh channel. It should be the numbers 1,2,3,..10

3.implement the same function using walk to determine whether t1 and t2 store the same values

4. Test the same function

Same(tree.New(1), tree.New(1)) should return, and
Same(tree.New(1), tree.New(2)) should return false

*/

package main

import "golang.org/x/tour/tree"

//walk walks the tree t sending all values from the tree to channel ch
func Walk(t *tree.Tree, ch chan int){
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value;
	Walk(t.Right, ch)
}

//same deltermines whther the trees t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree) bool {
	v1 := make(chan int)
	v2 := make(chan int)

	go Walk(t1, v1)
	go Walk(t2, v2)
	for i:=0; i<10; i++{
		value1 :=<- v1
		value2 :=<- v2
		if value1 != value2{
			return false
		}


	}
	return true
}


func main(){
	t1 := tree.New(1)
	ch := make(chan int, 10)
	Walk(t1, ch)
	for i:=0; i<10; i++{
		v:=<-ch
		println(v)
	}
	res1 :=  Same(tree.New(1), tree.New(1))
	res2 :=  Same(tree.New(1), tree.New(2))

	println(res1)
	println(res2)
}

