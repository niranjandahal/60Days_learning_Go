package main

import (
	"fmt"
)

func main() {

  number := make(chan int)
  message := make(chan string)

  go channelNumber(number)
  go channelMessage(message)

  select {
    case firstChannel := <-number:
      fmt.Println("Channel Data:", firstChannel)

    case secondChannel := <-message:
      fmt.Println("Channel Data:", secondChannel)
	
    default:
      fmt.Println("Wait!! Channels are not ready for execution")
  }

}



func channelMessage(message chan string) {

//   time.Sleep(2 * time.Second)
 
  message <- "Learning Go Select"
}

func channelNumber(number chan int) {
  
//   time.Sleep(2 * time.Second)
	
	number <- 15
}