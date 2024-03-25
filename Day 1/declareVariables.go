package main

import "fmt"

func main(){
	var smsSendingLimit int = 4
	var costPerSMS float32 = 0.6
	var hasPermission bool = true

	if hasPermission && smsSendingLimit > 0 {
    costPerSMS = costPerSMS * float32(smsSendingLimit) 
  }

	fmt.Printf("Total Cost Per Person: %f\n", costPerSMS)

}