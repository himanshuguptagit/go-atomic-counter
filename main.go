package main

import (
	"counter/counter"
	"fmt"
)

func main() {

	c := counter.NewCounter()

	n := 100000
	done := make(chan bool, n)

	for i := 0; i < n; i++ {
		go func() {
			c.Increment()
			done <- true
		}()
	}

	for i := 0; i < n; i++ {
		<-done
	}

	fmt.Println(c.GetCount())
}
