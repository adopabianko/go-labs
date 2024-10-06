// Buffered Channels
// url : https://www.meetgor.com/golang-channels/#:~:text=the%20second%20message.-,Buffered%20Channels,-In%20Go%2C%20you
package main

import (
	"fmt"
	"sync"
)

func main() {
	buffchan := make(chan int, 2)

	wg := sync.WaitGroup{}
	wg.Add(2)

	for i := 1; i <= 2; i++ {
		go func(n int) {
			buffchan <- n
			wg.Done()
		}(i)
	}

	wg.Wait()
	close(buffchan)

	for c := range buffchan {
		fmt.Println(c)
	}
}
