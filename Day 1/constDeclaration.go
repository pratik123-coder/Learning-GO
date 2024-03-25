package main

import "fmt"

func main() {

	//Constants must be known / value of the constant must be known before compilling
  const firstName = "Pratik"
	const lastName = "Mohanty"
	const fullName = firstName + " " + lastName

	fmt.Println("Your full name is", fullName)
}