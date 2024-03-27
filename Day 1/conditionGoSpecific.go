package main

import "fmt"

func main() {

	//rather than doing this
  length := getLength(email)
	if length < 1 {
		fmt.Printf("Email is invalid")
	}

	//we can do this
	if length:=getLength(email); length < 1 {
		fmt.Printf("Email is invalid")
	}

	/*reduces the scope of the variable and makes it 
	accessible only withing the if block*/
}