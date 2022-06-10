// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

/*
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel1 <- "I'll print every 1ms"
			time.Sleep(time.Millisecond * 1)

		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			channel2 <- "I'll print every 1s"
			time.Sleep(time.Second * 1)

		}
	}()

	for i := 0; i < 5; i++ {
		select {
		case message1 := <-channel1:
			fmt.Println(message1)
		case message2 := <-channel2:
			fmt.Println(message2)
		default:
			fmt.Println("No channel is ready")
		}
	}
}

/*
Output

No channel is ready
No channel is ready
I'll print every 1ms
I'll print every 1s
No channel is ready
*/
