package main

func worker(job chan int, result chan int) {
	for {
		a := <-job
		result <- a * a
	}
}

func Test() {

}
