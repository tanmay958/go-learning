package main

import "fmt"

func main() {
	var username string = "Tanmay"
	fmt.Printf("Variable is of type %T \n" , username) 

	var isLoggedIn bool = true
	fmt.Printf("Variable is of type %T \n" , isLoggedIn) 

	var smallVal uint8 = 255
	fmt.Printf("Variable is of type %T \n" , smallVal) 

	var flotVal float32 =  24.010484005
	fmt.Println(flotVal) 
	fmt.Printf("Variable is of type %T\n",flotVal)

	// default values
	var anotherVaribale string  
	fmt.Println(anotherVaribale)


	// implicit type
	var userJob = "Worker"
	fmt.Println(userJob)

	// no var type syntax
	gameName := "hye"
	fmt.Println(gameName) 

	// cant use the := syntax outside the main 

}