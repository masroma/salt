package main

import "fmt"

func main(){
	const lastName = "wick"
	fmt.Println("nice to meet you", lastName, "!\n")
	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")
}