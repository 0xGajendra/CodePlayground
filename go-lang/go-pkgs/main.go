package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
)

func main(){
	out, err := exec.Command("git", "status").Output()
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(out))
	fmt.Println("envv\n\n\n")
	num:=os.Getenv("TEST")
	fmt.Print(num)
}