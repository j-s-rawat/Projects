package main

import (
	"fmt"
	"sync"
	"time"
)

func usingMutex() {
	var x = 0

	var m sync.Mutex
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		/* 		go func(int, *sync.WaitGroup) {
			m.Lock()
			x = x + 1
			m.Unlock()
			w.Done()
		}(x, &w) */

		go func() {
			m.Lock()
			x = x + 1
			m.Unlock()
			w.Done()
		}()
	}
	w.Wait()
	fmt.Println(x)
}

func usingChannel() {
	var x = 0
	//******//
	//Very important to make it a buffered channel
	//*****//
	var c = make(chan struct{}, 1)
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go func() {
			c <- struct{}{}
			x = x + 1
			<-c
			w.Done()
		}()
	}
	w.Wait()
	fmt.Println(x)
}

func main() {
	usingMutex()
	usingChannel()
	//test()
	//test2()

}

// Basic sends and receives on channels are blocking.
// However, we can use `select` with a `default` clause to
// implement _non-blocking_ sends, receives, and even
// non-blocking multi-way `select`s.

func test() {
	messages := make(chan string)
	signals := make(chan bool)

	// Here's a non-blocking receive. If a value is
	// available on `messages` then `select` will take
	// the `<-messages` `case` with that value. If not
	// it will immediately take the `default` case.
	go func() {
		for {
			select {
			case msg := <-messages:
				fmt.Println("received message", msg)
				signals <- true
			default:
				fmt.Println("no message received")
				time.Sleep(1 * time.Microsecond)
			}
		}
	}()

	// A non-blocking send works similarly. Here `msg`
	// cannot be sent to the `messages` channel, because
	// the channel has no buffer and there is no receiver.
	// Therefore the `default` case is selected.
	msg := "hi"
	go func() {
		for {
			//this works but select doesn't why?
			messages <- msg
			/* select {
			case messages <- msg:
				fmt.Println("sent message", msg)
			default:
				fmt.Println("no message sent")
				time.Sleep(1 * time.Microsecond)
			} */

		}
	}()

	// We can use multiple `case`s above the `default`
	// clause to implement a multi-way non-blocking
	// select. Here we attempt non-blocking receives
	// on both `messages` and `signals`.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
	<-signals
}

//second example

func ping1(c chan string) {
	for i := 0; ; i++ {
		c <- "Ping on channel1"
	}
}

func ping2(c chan string) {
	for i := 0; ; i++ {
		c <- "Ping on channel2"
	}
}

func test2() {
	channel1 := make(chan string)
	channel2 := make(chan string)

	go ping1(channel1)
	go ping2(channel2)

	go func() {
		for {
			select {
			case msg1 := <-channel1:
				fmt.Println("Received", msg1)
			case msg2 := <-channel2:
				fmt.Println("Received", msg2)

			default:
				fmt.Println("Nothing ready, moving on.")
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}
