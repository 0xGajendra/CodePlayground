package main

import (
	"fmt"
	"os"
)

//exported struct (other packages can access it)
type User struct {
	Name string //exported field
	email string //private field
}

//unexported function (other packages cannot access it)
func validate(u User) bool {
	return u.Name != ""
}

// defer — cleanup guaranteed

//defer schedules a function call to be run when  surrounding fucntion return no matter what it returns. It is often used for cleanup tasks, such as closing files or releasing resources, ensuring that they are executed even if an error occurs or a return statement is reached.

func readFile(path string){
	f, _ := os.Open(path)
	defer f.Close() // ensures that the file is closed when readFile returns

	    // do work with f...
    // Close() is called automatically at the end of the function, even if an error occurs or a return statement is reached.

	//it is lifo (last in, first out) order, so if multiple defer statements are present, they will be executed in reverse order of their appearance in the code.
	defer fmt.Println("third")
defer fmt.Println("second")
defer fmt.Println("first")
}

//error handling
//In Go, error handling is typically done using the built-in error type. Functions that can encounter an error usually return an additional value of type error, which can be checked by the caller to determine if an error occurred.

func divide(a, b float64) (float64, error) {
	if(b==0){
		return 0, fmt.Errorf("Divide by zero error")

	}

	return a/b, nil
}

func main() {
	//variable declaration

	// := is a shorthand for declaring and initializing a variable in one step. It can only be used within a function and cannot be used for package-level variables.
	// var Name string = "John Doe" //explicit type declaration
	var name string = "John Doe"
	
	name = "gajendra"
	
	age:=22
	
	x, y := 10, 20
	
	x,y = y, x //swap values of x and y
	println("Hello, World!")
	fmt.Println("User Name:", name)
	fmt.Println("User Age:", age)
	fmt.Println("Swapped Values: x =", x, ", y =", y)

	readFile("/home/kenzo/Desktop/code/CodePlayground/go-lang/main.go")

	result, err := divide(1,0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}