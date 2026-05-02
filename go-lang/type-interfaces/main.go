package main

import "fmt"

func do (i interface{}) {
    fmt.Println(i)
	switch v:= i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
		
	}
}

func main(){
	do(21)
	do("hello")
	do(true)
}