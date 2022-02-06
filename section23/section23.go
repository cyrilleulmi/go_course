package section23

import (
	"fmt"
	"sync"
	"time"
)

func ex1() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func ex2() {
	cs := make(chan int)

	go func() {
		cs <- 42
	}()
	fmt.Println(<-cs)

	fmt.Printf("------\n")
	fmt.Printf("cs\t%T\n", cs)
}

func ex3() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}

		close(c)
	}()

	return c
}

func receive(c <-chan int) {
	for v := range c {
		println(v)
	}
}

func ex4() {
	q := make(chan int)
	c := gen_ex4(q)

	receive_ex4(c, q)
	close(q)

	fmt.Println("about to exit")
}

func receive_ex4(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			println(v)
		case <-q:
			return
		}
	}
}

func gen_ex4(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		q <- 0
		close(c)
	}()

	return c
}

func ex5() {
	c := make(chan int)

	go func() {
		c <- 12
	}()
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)
	time.Sleep(time.Second)
	v, ok = <-c
	fmt.Println(v, ok)
	v, ok = <-c
	fmt.Println(v, ok)
}

func ex6() {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		println(v)
	}
}

func ex7a() {
	c := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				for i := 0; i < 10; i++ {
					c <- i
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(c)
	}()

	for v := range c {
		println(v)
	}
}

func ex7b() {
	channels := []chan int{}

	channels = fanOut(channels)
	fannedInChannel := fanIn(channels)

	for v := range fannedInChannel {
		fmt.Println("received", v, "from fanned in channel")
	}
	fmt.Println("done")
}

func fanOut(channels []chan int) []chan int {
	for i := 0; i < 100; i++ {
		temporaryChannel := make(chan int)
		channels = append(channels, temporaryChannel)
		go sendValuesToChannel(temporaryChannel, i)
	}
	return channels
}

func sendValuesToChannel(c chan int, chanNumber int) {
	for j := 0; j < 9; j++ {
		fmt.Println("writing into temporary channel", chanNumber, "with value", j)
		c <- (j + 10*chanNumber)
	}
	close(c)
}

func fanIn(channels []chan int) chan int {
	resultingChannel := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for _, c := range channels {
			wg.Add(1)
			go handleChannelInput(resultingChannel, wg, c)
		}
		wg.Wait()
		close(resultingChannel)
	}()

	return resultingChannel
}

func handleChannelInput(resultingChannel chan int, wg sync.WaitGroup, funcChannel chan int) {
	for {
		v, ok := <-funcChannel
		if ok {
			resultingChannel <- v
		} else {
			wg.Done()
			break
		}
	}
}
