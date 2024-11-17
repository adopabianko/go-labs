// Select Statement for Channels
// url : https://www.meetgor.com/golang-channels/#:~:text=Below%20is%20an%20example%20to%20test%20which%20url%20or%20a%20web%20server%20is%20responding%20first%20to%20the%20request.
package main

import (
	"fmt"
	"net/http"
	"sync"
)

func pingGoogle(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, _ := http.Get("http://google.com")
	c <- res.Status
}

func pingDuckDuckGo(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, _ := http.Get("http://duckduckgo.com")
	c <- res.Status
}

func pingBraveSearch(c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	res, _ := http.Get("http://search.brave.com")
	c <- res.Status
}

func main() {
	gogChan := make(chan string)
	ddgChan := make(chan string)
	braveChan := make(chan string)

	var wg sync.WaitGroup
	wg.Add(3)

	go pingDuckDuckGo(ddgChan, &wg)
	go pingGoogle(gogChan, &wg)
	go pingBraveSearch(braveChan, &wg)

	openChannels := 3

	go func() {
		wg.Wait()
		close(gogChan)
		close(ddgChan)
		close(braveChan)
	}()

	for openChannels > 0 {
		select {
		case msg1, ok := <-gogChan:
			if !ok {
				openChannels--
			} else {
				fmt.Println("Google responded: ", msg1)
			}
		case msg2, ok := <-ddgChan:
			if !ok {
				openChannels--
			} else {
				fmt.Println("DuckDuckGo responded: ", msg2)
			}
		case msg3, ok := <-braveChan:
			if !ok {
				openChannels--
			} else {
				fmt.Println("Brave responded: ", msg3)
			}
		}
	}
}
