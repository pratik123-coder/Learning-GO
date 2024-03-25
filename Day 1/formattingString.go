package main

import "fmt"

func main() {
  const name = "Pratik Mohanty"
	const openRate = 30.6

	msg := fmt.Sprintf("Hi %s, your open rate is %.1f", name, openRate)

	fmt.Println(msg)
}