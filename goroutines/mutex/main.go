package main

import (
	"fmt"
	"runtime"
	"sync"
)

type counter struct {
	sync.Mutex
	val int
}

// tanpa c.lock() dan c.unlock() hasil tidak efektif, karena terdapat race condition
// untuk pengecekan bisa comment c.lock() dan c.unlock() lalu jalankan perintah go run -race main.go
func (c *counter) Add(int) {
	c.Lock()
	c.val++
	c.Unlock()
}

func (c *counter) value() int {
	return c.val
}

func main() {
	runtime.GOMAXPROCS(2)

	var (
		wg    sync.WaitGroup
		meter counter
	)

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				meter.Add(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.value())
}
