package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	content:= []byte("Hello niggers")

	err:= os.WriteFile("output.txt",content,0777)

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Done scene hai")

	readCont, err:=os.ReadFile("kya.txt")

	if err != nil{
		log.Fatal(err)
	}

	fmt.Printf("jsut read\n\n %v \n\nfrom file gooner", string(readCont))
}
