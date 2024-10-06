// Channel
// url : https://www.meetgor.com/golang-channels/#:~:text=various%20go%20routines.-,What%20are%20Channels,-A%20golang%20Channel
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	defer close(ch)

	go func() {
		message := "Hello, Gophers!"
		ch <- message
	}()

	msg := <-ch
	fmt.Println(msg)
}
