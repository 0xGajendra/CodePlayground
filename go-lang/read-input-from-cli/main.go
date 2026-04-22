package main

import (
	"errors"
	"fmt"
	"os"
)

func main()  {


	var path string;
	fmt.Println("Enter the fucking path:")
	fmt.Scan(&path);
	ans:=doesFileExists(path)

	if ans{
		println("file fucking exists")
	} else{
		println("file does not exixts nigga")
	}

	fmt.Println("path: ", path)
}

func doesFileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil{
		return true;
	}
	if errors.Is(err, os.ErrNotExist){
		return false
	}
	return false; //permission denied to read the file
}