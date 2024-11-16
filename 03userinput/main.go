package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to my world"
	println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating:")

	//comma ok syntax
	input , err :=  reader.ReadString('\n') 
	fmt.Println("thanks for the rating", input)
	fmt.Println("error message" ,err) 


}