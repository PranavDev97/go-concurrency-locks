package main

import (
	"fmt"
	"sync"
)

var counter = 0

//function to run as a goroutine.
//increments the value of counter by 1
func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock() //locks the access to below varible counter
	counter = counter + 1
	m.Unlock() //unlocks the access to the above variable

	wg.Done() //decrements the waitgroup counter by 1
}
func main() {

	var w sync.WaitGroup //variable to perform wait operation on goroutine
	var m sync.Mutex     //variable to lock and unlock the access to a variable

	for i := 0; i < 1000; i++ {
		w.Add(1) //adds value 1 to the waitgroup counter
		go worker(&w, &m)
	}

	w.Wait() //waits until the waitgroup counter value becomes 0
	fmt.Println("Value of counter", counter)
}
