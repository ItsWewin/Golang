package main

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- 1
	}

	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range int {

	}
}
