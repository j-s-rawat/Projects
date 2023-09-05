package main

/*

@author Arjun Dhar
@date 15th Sep 2021

*/

/*
USAGE:

go build -o sample.exe main.go
go run main.go

Install `Graphviz` @ https://graphviz.org/download/ and ensure part of System PATH
go tool pprof -png http://localhost:8082/debug/pprof/heap > heap.png
*/

import (
	//_ "net/http/pprof" //Init() <-- This is what plugs the route to <localhost>/debug/pprof/

	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/profile"
)

func main() {

	fmt.Println("[main]")

	//old http way of doin it
	//NewTest()

	//new way of doint it
	//to see the result on the browser
	//  go tool pprof -http=:8080 ./out/cpu.pprof

	fmt.Println("[Profiler started]")
	//mem_profile := profile.MemProfileRate(4096) //profile.MemProfile

	defer profile.Start(profile.CPUProfile, profile.ProfilePath("./out"), profile.NoShutdownHook).Stop()

	f()

}

var mu sync.Mutex

func f() {
	fmt.Println("[f()]")
	var s []string
	var x = 1

	// Periodically clear slcie ever intervalMin
	ticker := time.NewTicker(60 * 1000 * time.Millisecond) // NewTimer
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("[f->go] Timer fired : Clear slice["+strconv.Itoa(len(s))+"] @ ", t)
				clear(&s)
			}
		}
	}()

	for i := 0; i < 100000; i++ {
		appendTo(&s)
		x++
	}

	ticker.Stop()
	done <- true
}

func appendTo(s *[]string) {
	mu.Lock()
	*s = append(*s, "Hello how are you abcdefghijk")
	mu.Unlock()
	//time.Sleep(100 * time.Millisecond)
}

func clear(s *[]string) {
	fmt.Println("[clear]")
	mu.Lock()
	fmt.Println("[pre-clear] len=", len(*s))
	*s = (*s)[:0]
	*s = nil
	mu.Unlock()
	fmt.Println("[post-clear] len=", len(*s))
}

/*
PACKAGES:
* github.com/pkg/profile
	* https://pkg.go.dev/github.com/pkg/profile
* [Go /http/pprof](https://pkg.go.dev/net/http/pprof)

ARTICLES:
* [Gos pprof WebServer](https://github.com/golang/go/blob/master/src/net/http/pprof/pprof.go)
* [Find memory Leaks in Go](https://tusharsheth.medium.com/how-i-found-memory-leaks-in-the-golang-app-using-pprof-56e5d55363ba)
* [Memory leaks in Go](https://www.freecodecamp.org/news/how-i-investigated-memory-leaks-in-go-using-pprof-on-a-large-codebase-4bec4325e192/)
* [Go lang profiling](https://flaviocopes.com/golang-profiling/)
* [Go pprof trace examples](https://github.com/Raffo/go-pprof-trace-example)
*/
