package main

import (
	"fmt"
	"runtime"
	"sync"
)

func foo1() {
	fmt.Println("Hello from goroutine1!")
}
func bar1() {
	fmt.Println("Hello from goroutine2!")
	wg.Done()
}

type person struct {
	first string
	last  string
}

func (p *person) speak() { // method to type "person" with receiver type-> *person
	fmt.Printf("My name is %v %v.\n", p.first, p.last)
}

type human interface { // defining a human interface
	speak()
}

func saySomething(h human) {
	h.speak()
}

// defining wait group variable in package level scope
var wg sync.WaitGroup

func main() {
	fmt.Printf("\n--- Exercise 1: ---\n")
	fmt.Printf("Number of Goroutines: %v \n", runtime.NumGoroutine())

	wg.Add(1) // initiating one wait group

	foo1()
	go bar1() // letting another go routine handle the bar1() function call
	wg.Wait() // waiting for wg.done() signal from the previous goroutine

	//======================================================================================

	fmt.Printf("\n--- Exercise 2: ---\n")
	p1 := person{
		first: "slok",
		last:  "jha",
	}
	// show that
	// CAN pass a value of type *person into saySomething(...)
	saySomething(&p1)
	//CANNOT pass a value of type person into saySomething(...)
	// saySomething(p1) ... hence as saySomething() calls the method with receiver type *person...so it expects only *person

	//=======================================================================================

	fmt.Printf("\n--- Exercise 3,4,5: ---\n")
	incrementer := int64(0) // type conversion needed for "atomic.AddInt64()"

	wg.Add(10)
	var mu sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()

			hold := incrementer
			fmt.Println(incrementer)
			//runtime.Gosched() //  not needed if we use mutex lock()
			//hold
			hold++
			incrementer = hold

			// soln using sync/stomic pkg
			/*
				atomic.AddInt64(&incrementer, 1)
				r := atomic.LoadInt64(&incrementer)
				fmt.Println(r)
			*/

			mu.Unlock()

			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("final o/p: %v\n", incrementer)

	//=======================================================================================

	fmt.Printf("\n--- Exercise 6: ---\n")
	/*
		Create a program that prints out your OS and ARCH. Use the following commands to run it
			● go run
			● go build // creates the executable/package
			● go install // package has to be GOROOT to install
	*/
	fmt.Printf("OS: %v\n", runtime.GOOS)
	fmt.Printf("OS: %v\n", runtime.GOARCH)
}
