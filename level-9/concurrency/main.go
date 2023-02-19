package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup // package scope var

func main() {

	// waitgroup
	fmt.Println("OS\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Gooroutines\t\t", runtime.NumGoroutine()) // prints current number of goroutines

	wg.Add(1)
	go foo()
	bar()

	fmt.Println("CPUs\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())
	wg.Wait() // wait till wg.Done() is fired

	// !!!! Method sets ---- refer notes and docs !!!!
	// Goroutines and channels
	ch := make(chan int)
	go func() {
		ch <- doSomething(5) // sending the result through a channel
	}()
	fmt.Println(<-ch) // printing the result sent to a channel

	// Race Conditions and mutex: -----------------------------------------------------
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mu sync.Mutex // get rid of this mutex variable to simulate rwace condition in the code below.

	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock() // locks it if a routine is already accessing it in order to prevent race condition

			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()

			wg.Done()
		}()
		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter) // won't be equal to 100 as it should be

	// Atomic pkg -> refer docs ... an alternate to mutex ---------------------------------------------------------------

}

func doSomething(x int) int {
	return x * 5
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
