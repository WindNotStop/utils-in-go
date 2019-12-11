package sync

import (
	"fmt"
	"sync"
)

func Waitgroup() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		//do something
		fmt.Println("go")
	}()

	wg.Wait()
	fmt.Println("done")
}
