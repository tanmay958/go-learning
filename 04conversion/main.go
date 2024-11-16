package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("enter your rating:")
	reader:= bufio.NewReader(os.Stdin) 
	input,_  :=  reader.ReadString('\n') 
	numRating ,err :=  strconv.ParseFloat(strings.TrimSpace(input),64)
	if err!=nil {
		fmt.Println(err) 
	}else{ 
		fmt.Println("added 1 to the rating: ",  numRating+1) 
	}
}