# go-atomic-counter

Atomic Counter in Go

This project demonstrates a simple atomic counter implemented in Go. It utilizes goroutines and a mutex to ensure thread-safe increment operations on the counter.

Overview

The project consists of a main application and a counter package. The counter package provides the Counter type with methods to safely increment and retrieve the counter value. The main application spawns multiple goroutines to perform increments concurrently and prints the final counter value.

Files

main.go: Contains the main application logic.
counter/counter.go: Defines the Counter type and methods for incrementing and getting the counter value.
Getting Started

Prerequisites
Make sure you have Go installed. You can download and install Go from golang.org.

Running the Project
Clone the Repository (if applicable):
bash
Copy code
git clone <repository-url>
cd <repository-directory>
Navigate to the Project Directory:
bash
Copy code
cd path/to/project
Run the Application:
bash
Copy code
go run main.go
Code Explanation
main.go

go
Copy code
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
Initialization: Creates a new counter and a channel for synchronization.
Concurrency: Spawns 100,000 goroutines to increment the counter.
Synchronization: Waits for all goroutines to complete by reading from the done channel.
Output: Prints the final counter value.
counter/counter.go

go
Copy code
package counter

import "sync"

type Counter struct {
	count int
	mu    sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c Counter) GetCount() int {
	return c.count
}
Counter Struct: Defines a counter with a mutex to ensure safe concurrent access.
NewCounter(): Constructor function to create a new counter instance.
Increment(): Increments the counter value safely by locking and unlocking the mutex.
GetCount(): Returns the current count value.
Concurrency

The project demonstrates how to handle concurrency in Go using goroutines and mutexes. The Increment method ensures that the counter value is updated atomically, preventing race conditions.

License

This project is licensed under the MIT License. See the LICENSE file for more details.

Contributions

Feel free to fork the repository and make improvements. Pull requests are welcome!

Contact

For any questions or issues, please open an issue on the repository or contact the author at your-email@example.com.