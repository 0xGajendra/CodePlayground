package main

import "fmt"

func main(){
	var isValid bool;
	isValid = isGreaterThan10(15)
	fmt.Println(isValid)
}

func isGreaterThan10(num int) bool{
	return num > 10
}