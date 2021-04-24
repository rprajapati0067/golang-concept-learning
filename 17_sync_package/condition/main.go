package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	mu := sync.Mutex{}
	c := sync.NewCond(&mu)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		c.L.Lock()
		for len(sharedRsc) == 0 {
			c.Wait()
		}
		fmt.Println(sharedRsc["rsc1"])
		c.L.Unlock()
	}()

	c.L.Lock()
	sharedRsc["rsc1"] = "foo"
	c.Signal()
	c.L.Unlock()
	wg.Wait()
}
