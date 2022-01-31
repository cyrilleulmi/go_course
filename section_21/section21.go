package section21

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func ex1() {
	wg.Add(2)
	go subroutine1("ex1")
	go subroutine2("ex1")
	println("start waiting")
	wg.Wait()
	println("ex 1 finished")
}

func subroutine1(calledFrom string) {
	println("subroutine 1 called from", calledFrom)
	wg.Done()
}

func subroutine2(calledFrom string) {
	println("subroutine 2 called from", calledFrom)
	wg.Done()
}

func ex2() {
	myPerson := person{
		name: "Obi Wan",
	}

	pointerToPerson := &myPerson

	// saySomething(myPerson) doesn't work
	saySomething(pointerToPerson)
}

type person struct {
	name string
}

func (p *person) speak() {
	println("hello i'm", p.name)
}

type human interface {
	speak()
}

func saySomething(p human) {
	p.speak()
}

var counter int

func ex3() {
	counter = 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(i)
	}
	wg.Wait()
}

func increment(number int) {
	println("increment call", number, "read counter at ", counter)
	newCounter := counter
	runtime.Gosched()
	newCounter++
	println("increment call", number, "wrote counter with", newCounter)
	counter = newCounter
	wg.Done()
}

func ex4() {
	counter = 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go mutexIncrement(i)
	}
	wg.Wait()
}

var mu sync.Mutex

func mutexIncrement(number int) {
	mu.Lock()
	println("increment call", number, "read counter at ", counter)
	newCounter := counter
	newCounter++
	println("increment call", number, "wrote counter with", newCounter)
	counter = newCounter
	wg.Done()
	mu.Unlock()
}

func ex5() {
	counter = 0
	atomicCounter = 0
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go atomicIncrement(i)
	}
	wg.Wait()

	println("------------------------------------")

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go normalIncrement(i)
	}
	wg.Wait()
}

var atomicCounter int32

func normalIncrement(number int) {
	println("increment call", number, "read counter at ", counter)
	counter++
	println("increment call", number, "wrote counter with", counter)
	wg.Done()
}

func atomicIncrement(number int) {
	println("increment call", number, "read counter at ", atomicCounter)
	atomic.AddInt32(&atomicCounter, 1)
	println("increment call", number, "wrote counter with", atomicCounter)
	wg.Done()
}
