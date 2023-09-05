// _Closing_ a channel indicates that no more values
// will be sent on it. This can be useful to communicate
// completion to the channel's receivers.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func longTimedOperation() chan int32 {
	ch := make(chan int32)
	go func() {
		defer close(ch)
		//time.Sleep(time.Second * 5)
		ch <- rand.Int31n(300)
	}()
	//	go trun()
	return ch
}

// In this example we'll use a `jobs` channel to
// communicate work to be done from the `main()` goroutine
// to a worker goroutine. When we have no more jobs for
// the worker we'll `close` the `jobs` channel.
func main() {
	ch := longTimedOperation()
	i, more := <-ch
	fmt.Println(i, more) //more is true as the channel is not
	i, more = <-ch
	fmt.Println(i, more) //default value of channel is retured as nothing is added to the channel and more is flase
	jobs := make(chan int, 5)
	//done := make(chan bool)

	// Here's the worker goroutine. It repeatedly receives
	// from `jobs` with `j, more := <-jobs`. In this
	// special 2-value form of receive, the `more` value
	// will be `false` if `jobs` has been `close`d or all
	// values in the channel have already been received.
	// We use this to notify on `done` when we've worked
	// all our jobs.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				//done <- true
				return
			}
		}
	}()

	var wg sync.WaitGroup

	wg.Add(3)
	// This sends 3 jobs to the worker over the `jobs`
	// channel, then closes it.
	for j := 1; j <= 3; j++ {
		go func(wg *sync.WaitGroup, j int) {
			jobs <- j
			fmt.Println("sent job", j)

			time.Sleep(5 * time.Second)
			wg.Done()
		}(&wg, j)
	}
	wg.Wait()
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the
	// [synchronization](channel-synchronization) approach
	// we saw earlier.
	//<-done
}
