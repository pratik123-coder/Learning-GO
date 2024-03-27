package main

import "fmt"

func main() {
	messageLen :=10
	maxMessageLen := 20

	fmt.Println("Trying to send a message of length:", messageLen)

	if messageLen > maxMessageLen {
    fmt.Println("The message is too long!\n")
		fmt.Println("Hence the message can't be sent")
  } else {
    fmt.Println("The message is valid!\n")
		fmt.Println("Hence the message is sent")
  }

}